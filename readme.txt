==== social network ====

---- http handling ----
using "net/http" for routing & serving

---- postgres access ----
using "database/sql" for connection & queries
and "github.com/lib/pq" for the postgres driver

---- database tables/go structs ----
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(8) NOT NULL UNIQUE,
	chosen_name VARCHAR(128) NOT NULL,
	email VARCHAR(128) NOT NULL UNIQUE,
	password_hash TEXT NOT NULL
);
type User struct {
	ID		int64
	Username	string
	ChosenName	string
	Email		string
	PasswordHash	string
)