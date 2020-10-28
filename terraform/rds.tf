
resource "aws_db_instance" "main" {
  allocated_storage    = 20
  storage_type         = "gp2"
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t2.micro"
  name                 = "main"
  username             = "newarchitectures"
  password             = "newarchitectures"
  parameter_group_name = "default.mysql5.7"
}