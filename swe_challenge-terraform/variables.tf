variable "aws_region" {
  description = "Región de AWS donde se desplegarán los recursos"
  type        = string
  default     = "us-east-1"
}

variable "instance_type" {
  description = "Tipo de instancia de EC2"
  type        = string
  default     = "t2.micro"  # Gratis en Free Tier
}

variable "db_name" {
  description = "Nombre de la base de datos"
  type        = string
  default     = "stockdb"
}

variable "db_username" {
  description = "Nombre de usuario"
  type        = string
  default     = "pruebatec"
}

variable "db_password" {
  description = "Clave bd"
  default     = "XF-yM8ZYvdWF-7GuzVRf9Q"
}

variable "next_page_limit" {
  description = "Cantidad de páginas a traer desde la API externa"
  type        = number
  default     = 20
}