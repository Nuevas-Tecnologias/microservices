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

