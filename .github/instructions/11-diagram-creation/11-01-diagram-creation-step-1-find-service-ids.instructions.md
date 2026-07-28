---
applyTo: ".github/instructions/manual/**"
---

# 11.01 Diagram creation: Step 1 — Find Service IDs

## Step 1 — Find Service IDs

`etc/resources/aws/service-index.csv` maps service IDs to service names.
Use `grep` to search for the services you need.

```bash
# Format: id,service
grep -i "ec2"          etc/resources/aws/service-index.csv
grep -i "rds\|aurora"  etc/resources/aws/service-index.csv
grep -i "cloudfront"   etc/resources/aws/service-index.csv
```

Example output:
```
27,Amazon EC2
117,Amazon RDS
1178,Amazon CloudFront
```

---
