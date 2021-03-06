package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
)

// GenerateOID from given string or id.
func GenerateOID(ID string) (string, error) {
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return "", Error(codes.Aborted, "Unable to decode id", err)
	}

	return OID.String(), nil
}
