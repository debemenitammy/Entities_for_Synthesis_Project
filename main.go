package main

// Struct for `cohortplayers`
type CohortsPlayers struct {
	cohortplayerid int32
	playerid int32
	cohortid int32
	types string
}

// Struct for `cohorts``
type Cohort struct {
  cohortid int32
  name string
  date string
  at_sku int32
  at_id string
  status string
  zoom string
  game string
  level int8
  imported_students int8
  start_date string
  end_date string
  mapaccess int8
  record int8
  event_id string
}

// Struct for `cohort_setup_permissions`
type CohortSetupPermissions struct {
	cohortid int32
	simulationid int32
}

// Struct for `games`
type Games struct {
	gameid int32
	gametype string
	mapid int32
	name string
	datetime string
	token string
	rounds int32
	teams int32
	roundlength int32
	mtime float64
	currentround int32
	distance_weight float64
	number_weight float64
	longest_weight float64
	longest_hq_weight int32
	multimove int8
	templateid int32
	cohortid int32
	active int8
	current int8
	opts string
	stealcount int8
	record int8
	ranked int8
}

// Struct for `nots`
type Nots struct {
	notid int32
	playerid int32
	cohortid int32
	tag string
	subtag string
	timestamp string
}

// Struct for `players`
type Players struct {
	playerid int32
	cohortid int32
	name string
	email string
	age int32
	dob string
	username string
	password string
	admin int8
	order int8
	reset_token string
	password_sent string
	last_login string
	timezone string
	at_id string
	hashed_password string
	start_date string
	user_id string
}


//Struct for `simulations`
type Simulations struct {
	simulationid int32
	name string
	game_url string
	setup_url string
	env string
	configvar string
	field_mapping string
}

// Struct for `teamplayers`
type TeamPlayers struct {
	id int32
	teamid int32
	playerid int32
	active int8
}


// Struct for `teams`
type Teams struct {
  teamid int32
  gameid int32
  color string
  token string
  hassteal int8
  longest int32
}