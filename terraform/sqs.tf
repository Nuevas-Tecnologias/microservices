
resource "aws_sqs_queue" "process_tech_revision_format_command" {
  name                        = "process-tech-revision-format-command.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
}

resource "aws_sqs_queue" "tech_revision_ack" {
  name                        = "tech-revision-ack.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
}

resource "aws_sqs_queue" "save_tech_order_command" {
  name                        = "save-tech-order-command.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
}

resource "aws_sqs_queue" "assign_service_center_command" {
  name                        = "assign-service-center-command.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
}

resource "aws_sqs_queue" "create_service_register_bulk_command" {
  name                        = "create-service-register-bulk-command.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
}

