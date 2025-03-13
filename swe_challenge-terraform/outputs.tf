output "ec2_instance_public_ip" {
  description = "Dirección IP pública de la instancia EC2"
  value       = aws_instance.backend.public_ip
}

output "database_endpoint" {
  description = "Endpoint de la base de datos en AWS RDS"
  value       = aws_db_instance.cockroachdb.endpoint
}
