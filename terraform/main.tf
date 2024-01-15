
# Terraform State -- fill this as required
terraform {
    backend "s3" {
        region = "aws-region"
        key = "terraform.tfstate"
        bucket = "state-bucket-name"
    }
}

resource "aws_security_group" "security_group" {
    name = "app_security_group_name"
    description = "Allow web inbound traffic"
    vpc_id = "vpc-id"

    # Allow inbound traffic from port 80
    ingress {
        description = "HTTP"
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    # Allow outbound traffic
    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = "app security group"
    }
}


resource "aws_instance" "app_instance_name" {
    ami = var.ami_id
    instance_type = var.instance_type
    key_name = "app-ssh-key-pair-name"
    security_groups = [aws_security_group.chat_app_security_group.name]

    tags = {
        Name = "app-instance"
    }
}


resource "aws_s3_bucket" "app_bucket" {
    bucket = "app-bucket-name"
}


resource "aws_eip" "app_eip" {
    instance = aws_instance.chat_app_instance.id
}