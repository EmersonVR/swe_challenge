terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  required_version = ">= 1.3.0"
}

provider "aws" {
  region  = "us-east-1"  # Cambia según la región que prefieras
  profile = "default"    # Usa el perfil configurado con `aws configure`
}
