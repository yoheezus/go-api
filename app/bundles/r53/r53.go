package r53

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

type RecordSet struct {
	HostedZoneId    string
	RecordSetName   string
	RecordType      string // "Type": "SOA"|"A"|"TXT"|"NS"|"CNAME"|"MX"|"PTR"|"SRV"|"SPF"|"AAAA",
	ResourceRecords []string
}

// List record sets in a specificed hosted zone
func GetHostedZone(id string) *route53.HostedZone {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: credentials.NewSharedCredentials("", "dom-personal"),
	})
	if err != nil {
		println(err.Error())
	}
	svc := route53.New(sess)
	input := &route53.GetHostedZoneInput{Id: aws.String(id)}
	result, err := svc.GetHostedZone(input)
	if err != nil {
		println("Failed to get hosted zones: " + err.Error())
		return nil
	}
	return result.HostedZone
}

// List record sets in a specificed hosted zone
func GetRecordSets(id string) []*route53.ResourceRecordSet {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: credentials.NewSharedCredentials("", "dom-personal"),
	})
	if err != nil {
		println(err.Error())
	}

	svc := route53.New(sess)
	input := &route53.ListResourceRecordSetsInput{HostedZoneId: aws.String(id)}
	result, err := svc.ListResourceRecordSets(input)
	if err != nil {
		println("Failed to get hosted zones: " + err.Error())
		return nil
	}
	return result.ResourceRecordSets
}

// Updates or inserts a record set
func UpdateRecordSet(rs RecordSet) string {
	println(rs.HostedZoneId)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: credentials.NewSharedCredentials("", "dom-personal"),
		LogLevel:    aws.LogLevel(aws.LogDebugWithHTTPBody),
	})
	if err != nil {
		println(err.Error())
	}

	svc := route53.New(sess)
	action := "UPSERT"
	var resRecs []*route53.ResourceRecord
	for _, rec := range rs.ResourceRecords {
		resRecs = append(resRecs, &route53.ResourceRecord{Value: aws.String(rec)})
	}
	var ttl int64
	ttl = 60
	input := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: &rs.HostedZoneId,
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				&route53.Change{
					Action: &action,
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name:            &rs.RecordSetName,
						Type:            &rs.RecordType,
						TTL:             &ttl,
						ResourceRecords: resRecs,
					},
				},
			},
		},
	}
	println("Calling ChangeResourceRecordSets")
	res, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		println("ERROR: " + err.Error())
		return "ERROR: " + err.Error()
	}
	return *res.ChangeInfo.Status
}
