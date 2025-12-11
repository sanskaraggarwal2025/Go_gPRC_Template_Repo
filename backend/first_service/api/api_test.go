package api

import (
	"context"
	"testing"

	"first_service/dao"
	"first_service/proto"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func TestSayHello(t *testing.T) {
	// 1. Setup in-memory database
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	// 2. Run migrations
	if err := dao.Migrate(db); err != nil {
		t.Fatalf("failed to migrate db: %v", err)
	}

	// 3. Initialize DAO
	sqliteDAO := dao.NewSQLiteDAO(db)

	// 4. Initialize ServiceServer
	server := &ServiceServer{
		Dao: sqliteDAO,
	}

	// 5. Call SayHello
	req := &proto.HelloRequest{Naam: "TestUser"}
	resp, err := server.SayHello(context.Background(), req)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}

	// 6. Verify Response
	expectedMessage := "Hello TestUser"
	if resp.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, resp.Message)
	}

	// 7. Verify Database Insertion
	var messages []dao.Message
	err = db.Select(&messages, "SELECT * FROM first_service_message")
	if err != nil {
		t.Fatalf("failed to query messages: %v", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message in db, got %d", len(messages))
	}

	if messages[0].Message != "TestUser" {
		t.Errorf("expected db message %q, got %q", "TestUser", messages[0].Message)
	}
}
