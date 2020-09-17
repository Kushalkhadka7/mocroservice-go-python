package util

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateOID from given string or id.
func GenerateOID(ID string) (string,error){
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return "",err
	}

	return OID.String(),nil
}