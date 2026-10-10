resource "aws_kms_replica_external_key" "test" {
  description             = var.rName
  enabled                 = true
  primary_key_arn         = aws_kms_external_key.test.arn
  deletion_window_in_days = 7

  key_material_base64 = "Uhrl02ropsyLonnvC7kGZgFEJemYoo8wMmZtq49g/PW="

{{- template "tags" . }}
}

# ACCEPTANCE TESTING ONLY -- NEVER EXPOSE YOUR KEY MATERIAL
resource "aws_kms_external_key" "test" {
  provider = awsalternate

  description  = "${var.rName}-source"
  multi_region = true
  enabled      = true

  key_material_base64 = "Rsdy06lxkevFpeoqM5tLRiQYBpdCly5mHrLxk02t/WQ="

  deletion_window_in_days = 7
}
