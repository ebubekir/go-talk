package main

import "os"

type EnvType string

const (
	Prod        EnvType = "prod"
	Qa          EnvType = "qa"
	Development EnvType = "development"
)

type Environment struct {
	EnvironmentType          EnvType
	AppUrl                   string
	ServiceName              string
	MongoDbConnectionString  string
	FirebaseProjectId        string
	FirebaseConnectionString string
}

func InitializeEnvironment() *Environment {
	environmentType := Development
	switch getValue("ENV", string(Development)) {
	case string(Qa):
		environmentType = Qa
	case string(Prod):
		environmentType = Prod
	}

	env := &Environment{
		EnvironmentType:         environmentType,
		ServiceName:             getValue("SERVICE_NAME", "api-gotalk"),
		AppUrl:                  getValue("APP_URL", "http://localhost:3000"),
		MongoDbConnectionString: getValue("MONGODB_CONNECTION_STRING", "mongodb://mongo:1234@localhost:27017/?tls=false"),
		FirebaseConnectionString: getValue("FIREBASE_CONNECTION_STRING", `{
					  "type": "service_account",
					  "project_id": "gotalk-85909",
					  "private_key_id": "5d986ed3d3564aaa90c25093a7c9a85eab0c2bba",
					  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDA849sB7zn4wtV\nLSaWrumNLyepS6mdaPkvDWxtNlxbmSf/crDIewwiynjN2iWi8OsJ5ig5hwSpWotm\nokk+Bdf5WSqb4rH7zJhYe7svt0fyMzmo77HIyGFOeYkYAuhMsNEzNSpTpaikDTHY\nzHAskQftWJ6KCHA6W82JIVW0MVzhIqrtworhuyXPd5T5DUxsHnugKOfuEIOLpTyH\nPfSR+caX+4Sa3xl4juZi/i72iBSJBxJLenh6a9OmahJZf37JEVesYtk1qn7WlVg7\nM+bWKzjW+fnxsmDpy3rwhPosbWNHfgO4lfsDgK2JPyISJWxeSfwNBmn6inUrzQ7H\nRDvD+8jzAgMBAAECggEAOLsj7ypzQ2bBHUESOHmjF6zGp4MkkrCbm1cCCzZRf2kP\nlo7dJYTwM4Z+cF/0cu0M3jM6nndxSm3h0MJkcIT9VEYAPicwF423OUTf646i67Zd\n/KrFBfjMi2s2gMXSEUJnr/uwvzlU1S8/+bNaQ/A8eW915bXHcZEuZGRVs45T9ale\nAPa0P61LEsFl77U9VGp91ebEsqh8k4BpCvLMLMl5NpGGoXxhF8VrrNZOSjVWwnUY\nzVhZ2XlqtGXSV2VXsIygwKHuXCdvH0ldCA7z/DF+vbR5fy89nrvIWNaI3kVqdRIa\n/3PmI/HgEHxkvNb/vCShXRfr2n9pFiKIjK8vPakl+QKBgQDlNs0H/8MrI26zUqnI\nfhibwDxeEhtAXKQk9P1Ypx+Zi1DpzGwy3HlXXc9w4O+JkpG5WQx+Z+x3d94XZiD/\nNM67qmixXiCWhzYq04KckcAyOkWZcwoDAo0OEV2XSpdj7qgZ10DkNGbsrs8xaywX\nbD27vIvUHi92UDU4hNjlQyUfJQKBgQDXf+u0I/2ny726nZCUXl5GAlNcarWRdmNO\nIWS3t2N4IDMf8PAmT+XYHTdqk9313+NZQo2N2aUaINsL2NeONUAv+qk4TTmlQuqv\nYrlTyoQz8J0O2/4PuzJfylmmSFSZdBBGucu8HD/zA81w6ICBSnjNEMixK50fmXnN\na/C2gco4NwKBgQCNGwxjZT8n2ls2x6e/xmit0U0YeDsQhzeBjNQ99DxO0OYR2Aev\n0+xbLWQb0E2GOpW9LaW0V5PKBB/T9cpQcZjnDMQAlLqpEDn3aVgZvNw9z4OzMI+0\nRKjDRUuBbKkAGxafOdU3506JXCAvAxQUo0zpuuu1vJNpWX05+wZvNMOwhQKBgQC+\nniZt15Adhniyw4EJ3Fdjcdcu3izxGFlK2PrwDsVrkn/mdwbVvMLAYUeNfHJPdNTz\nNY1kten2rK1VU1+IKM44Im7goF6nMgPJU3g/B9nc367tX+bhH2K1nJWkIkLC4gkr\npljyccKXQPvOLbrNooQsT/ZV0RBQzT8SV7I6nZ+0DQKBgDV+sHGNIuhQwg0FIj4E\n4j0hxgzW8z+4uCn+gnO5lvdD2v6zAKE4I8HFxc7s/WryYRPztE9C+MBwq49g3d11\nc3sfGoOlwHmYIJHY710fStgb5ODBic8f1mwfKjqB/o6CFwEEKK5aKIa/YvawJnjv\nRl6E+qF6uVEL0jME/i/2XQBq\n-----END PRIVATE KEY-----\n",
					  "client_email": "firebase-adminsdk-fbsvc@gotalk-85909.iam.gserviceaccount.com",
					  "client_id": "100139951561413883532",
					  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
					  "token_uri": "https://oauth2.googleapis.com/token",
					  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
					  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-fbsvc%40gotalk-85909.iam.gserviceaccount.com",
					  "universe_domain": "googleapis.com"
					}
		`),
		FirebaseProjectId: getValue("FIREBASE_PROJECT_ID", "gotalk-85909"),
	}
	return env
}

func getValue(key string, defaultValue string) string {
	value, hasValue := os.LookupEnv(key)
	if hasValue {
		return value
	}

	return defaultValue
}
