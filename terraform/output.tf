
output "ec2_ip" {
  description = "The Elastic IP address of the instance"
  value = aws_eip.chat_app_eip.public_ip
}