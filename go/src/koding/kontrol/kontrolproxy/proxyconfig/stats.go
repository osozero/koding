package proxyconfig

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"strconv"
	"time"
)

type DomainStat struct {
	Id           bson.ObjectId  `bson:"_id" json:"-"`
	Domainname   string         `bson:"domainname" json:"domainname"`
	RequestsHour map[string]int `bson:"requesthour", json:"requesthour"`
}

type ProxyStat struct {
	Id           bson.ObjectId  `bson:"_id" json:"-"`
	Proxyname    string         `bson:"proxyname" json:"proxyname"`
	Country      map[string]int `bson:"country", json:"country"`
	RequestsHour map[string]int `bson:"requesthour", json:"requesthour"`
}

func NewDomainStat(name string) *DomainStat {
	return &DomainStat{
		Id:           bson.NewObjectId(),
		Domainname:   name,
		RequestsHour: make(map[string]int),
	}
}

func NewProxyStat(name string) *ProxyStat {
	return &ProxyStat{
		Id:           bson.NewObjectId(),
		Proxyname:    name,
		Country:      make(map[string]int),
		RequestsHour: make(map[string]int),
	}
}

func (p *ProxyConfiguration) AddDomainStat(domainname string) error {
	domainStat, err := p.GetDomainStat(domainname)
	if err != nil {
		return err
	}

	if domainStat.Domainname == "" {
		domainStat = *NewDomainStat(domainname)
	}

	nowHour := strconv.Itoa(time.Now().Hour()) + ":00"
	_, ok := domainStat.RequestsHour[nowHour]
	if !ok {
		domainStat.RequestsHour[nowHour] = 1
	} else {
		domainStat.RequestsHour[nowHour]++
	}

	_, err = p.Collection["domainstats"].Upsert(bson.M{"domainname": domainname}, domainStat)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProxyConfiguration) DeleteDomainStat(domainname string) error {
	err := p.Collection["domainstats"].Remove(bson.M{"domainname": domainname})
	if err != nil {
		return err
	}
	return nil
}

func (p *ProxyConfiguration) GetDomainStat(domainname string) (DomainStat, error) {
	domainstat := DomainStat{}
	err := p.Collection["domainstats"].Find(bson.M{"domainname": domainname}).One(&domainstat)
	if err != nil {
		if err.Error() == "not found" {
			return domainstat, nil //return empty struct
		}
		return domainstat, fmt.Errorf("no stat for domain %s exist (error %s).", domainname, err.Error())
	}

	return domainstat, nil
}

func (p *ProxyConfiguration) GetDomainStats() []DomainStat {
	domainstat := DomainStat{}
	domainstats := make([]DomainStat, 0)
	iter := p.Collection["domainstats"].Find(nil).Iter()
	for iter.Next(&domainstat) {
		domainstats = append(domainstats, domainstat)
	}

	return domainstats
}

func (p *ProxyConfiguration) AddProxyStat(proxyname, country string) error {
	proxyStat, err := p.GetProxyStat(proxyname)
	if err != nil {
		return err
	}

	if proxyStat.Proxyname == "" {
		proxyStat = *NewProxyStat(proxyname)
	}

	nowHour := strconv.Itoa(time.Now().Hour()) + ":00"
	_, ok := proxyStat.RequestsHour[nowHour]
	if !ok {
		proxyStat.RequestsHour[nowHour] = 1
	} else {
		proxyStat.RequestsHour[nowHour]++
	}

	if country != "" {
		_, ok = proxyStat.Country[country]
		if !ok {
			proxyStat.Country[country] = 1
		} else {
			proxyStat.Country[country]++
		}
	}

	_, err = p.Collection["proxystats"].Upsert(bson.M{"proxyname": proxyname}, proxyStat)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProxyConfiguration) DeleteProxyStat(proxyname string) error {
	err := p.Collection["proxystats"].Remove(bson.M{"proxyname": proxyname})
	if err != nil {
		return err
	}
	return nil
}

func (p *ProxyConfiguration) GetProxyStat(proxyname string) (ProxyStat, error) {
	proxystat := ProxyStat{}
	err := p.Collection["proxystats"].Find(bson.M{"proxyname": proxyname}).One(&proxystat)
	if err != nil {
		if err.Error() == "not found" {
			return proxystat, nil //return empty struct
		}
		return proxystat, fmt.Errorf("no stat for proxy %s exist (error %s).", proxyname, err.Error())
	}

	return proxystat, nil
}

func (p *ProxyConfiguration) GetProxyStats() []ProxyStat {
	proxystat := ProxyStat{}
	proxystats := make([]ProxyStat, 0)
	iter := p.Collection["proxystats"].Find(nil).Iter()
	for iter.Next(&proxystat) {
		proxystats = append(proxystats, proxystat)
	}

	return proxystats
}
