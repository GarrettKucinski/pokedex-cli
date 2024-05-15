package main

func mapCommand(cfg *config) (mapError error) {
	cfg.mapClient.DisplayNextLocationList()

	return
}

func mapbCommand(cfg *config) (mapbError error) {
	cfg.mapClient.DisplayPrevLocationList()

	return
}
