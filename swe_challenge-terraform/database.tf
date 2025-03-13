resource "aws_db_instance" "cockroachdb" {
  allocated_storage    = 20
  engine               = "postgres"
  engine_version       = "12.19"
  instance_class       = "db.t3.micro"
  identifier           = var.db_name
  username            = var.db_username
  password            = var.db_password
  parameter_group_name = "default.postgres12"
  publicly_accessible  = true
  skip_final_snapshot  = true
}
