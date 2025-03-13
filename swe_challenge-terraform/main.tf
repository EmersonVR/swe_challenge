resource "aws_instance" "backend" {
  ami           = "ami-045602374a1982480" # Amazon Linux 2
  instance_type = var.instance_type
  key_name      = "my-key1"  # Usa el par de claves que creaste en AWS

  tags = {
    Name = "Backend-Golang"
  }

  user_data = <<-EOF
              #!/bin/bash
              sudo yum update -y
              sudo amazon-linux-extras enable go1.19
              sudo yum install -y golang git

              # Clonar el código del repositorio
              git clone https://github.com/EmersonVR/swe_challenge.git /home/ec2-user/swe_challenge
              cd /home/ec2-user/swe_challenge/cmd

              # Construir y mover el ejecutable
              go build -o app
              sudo mv app /usr/local/bin/backend-app

              # Crear un servicio para que el backend siempre corra
              sudo tee /etc/systemd/system/backend.service <<EOL
              [Unit]
              Description=Backend Golang API
              After=network.target

              [Service]
              Type=simple
              ExecStart=/usr/local/bin/backend-app
              Restart=always
              User=ec2-user
              WorkingDirectory=/home/ec2-user/swe_challenge/cmd

              [Install]
              WantedBy=multi-user.target
              EOL

              # Recargar servicios y arrancar la API
              sudo systemctl daemon-reload
              sudo systemctl enable backend.service
              sudo systemctl start backend.service
              EOF
}
