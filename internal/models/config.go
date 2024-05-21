package models

type Config struct {
    DBHost     string `json:"db_host"`
    DBPort     string `json:"db_port"`
    DBUser     string `json:"db_user"`
    DBPassword string `json:"db_password"`
    DBName     string `json:"db_name"`
}
