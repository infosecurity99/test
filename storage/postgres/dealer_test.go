package postgres

/*
func TestDealer_AddSum(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	sum := 100
	dealerID := "1cfd84e6-72cb-4135-a802-85d10e4183ea"
	err = pgStore.Dealer().AddSum(context.Background(), sum)
	if err != nil {
		t.Fatalf("error while adding sum to dealer: %v", err)
	}
pgStore.Dealer()
	// Query the updated sum directly from the database after adding the sum.
	var updatedSum int
	err = pgStore.db.QueryRowContext(context.Background(), "SELECT sum FROM dealer WHERE id = $1", dealerID).Scan(&updatedSum)
	if err != nil {
		t.Fatalf("error while querying dealer sum: %v", err)
	}

	expectedSum := 100
	if updatedSum != expectedSum {
		t.Fatalf("expected dealer sum %d, got %d", expectedSum, updatedSum)
	}
}
*/
