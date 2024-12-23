package repositories

import (
	"context"
	"fmt"
	"gateway/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conversationsCollection *mongo.Collection

func InitMongoDB(uri, database, collection string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Initialize collection
	conversationsCollection = client.Database(database).Collection(collection)
	fmt.Println("MongoDB connection established")
	return client, nil
}

// SaveMessage confere se a última mensagem trocada faz menos de 24 horas.
// Se não, adiciona a mensagem na conversa ativa.
// Se sim, inicia uma nova conversa
func SaveMessage(userID string, newMessage models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Encontra a ultima conversa com o usuário
	var recentConversation models.Conversation
	err := conversationsCollection.FindOne(ctx, bson.M{"user_id": userID}, options.FindOne().SetSort(bson.M{"start_time": -1})).Decode(&recentConversation)
	log.Printf("FindOne query result: %+v, err: %v", recentConversation, err)

	// Determina se é necessário criar uma nova conversa
	if err == mongo.ErrNoDocuments || recentConversation.EndTime != nil && time.Since(*recentConversation.EndTime) > 24*time.Hour {
		// Nenhuma conversa recente ou se conversa nas últimas 24 horas
		// Fechando a última conversa se existir
		if err != mongo.ErrNoDocuments {
			_, err := conversationsCollection.UpdateOne(ctx, bson.M{"_id": recentConversation.ID}, bson.M{
				"$set": bson.M{"end_time": time.Now()},
			})
			if err != nil {
				return fmt.Errorf("failed to close conversation: %v", err)
			}
		}

		// Criando uma nova conversa
		newConversation := models.Conversation{
			UserID:    userID,
			StartTime: newMessage.Timestamp,
			Messages:  []models.Message{newMessage},
		}
		_, err := conversationsCollection.InsertOne(ctx, newConversation)
		if err != nil {
			return fmt.Errorf("failed to create new conversation: %v", err)
		}

		return nil
	}

	// Convertendo a string para um primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(recentConversation.ID)
	if err != nil {
		return fmt.Errorf("failed to convert string to ObjectID: %v", err)
	}

	// Adicionando a mensagem na conversa ativa
	_, err = conversationsCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$push": bson.M{"messages": newMessage},
	})
	if err != nil {
		return fmt.Errorf("failed to append message to conversation: %v", err)
	}

	return nil
}
