package http

import (
	"bytes"
	"encoding/json"
	cmodel "github.com/Cepave/common/model"
	"github.com/Cepave/query/g"
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Tag struct {
	StrategyId int
	Name       string
	Value      string
	CreateAt   string
	UpdateAt   string
}

/**
 * @function name:   func postByJson(rw http.ResponseWriter, req *http.Request, url string)
 * @description:     This function sends a POST request in JSON format.
 * @related issues:  OWL-171
 * @param:           rw http.ResponseWriter
 * @param:           req *http.Request
 * @param:           url string
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/12/2015
 * @last modified:   11/13/2015
 * @called by:       func queryInfo(rw http.ResponseWriter, req *http.Request)
 *                   func queryHistory(rw http.ResponseWriter, req *http.Request)
 */
func postByJSON(rw http.ResponseWriter, req *http.Request, url string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	s := buf.String()
	reqPost, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(s)))
	if err != nil {
		log.Println("Error =", err.Error())
	}
	reqPost.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqPost)
	if err != nil {
		log.Println("Error =", err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Write(body)
}

/**
 * @function name:   func queryInfo(rw http.ResponseWriter, req *http.Request)
 * @description:     This function handles /graph/info API request.
 * @related issues:  OWL-171
 * @param:           rw http.ResponseWriter
 * @param:           req *http.Request
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/12/2015
 * @last modified:   11/13/2015
 * @called by:       func configApiRoutes()
 */
func queryInfo(rw http.ResponseWriter, req *http.Request) {
	url := g.Config().Api.Query + "/graph/info"
	postByJSON(rw, req, url)
}

/**
 * @function name:   func queryHistory(rw http.ResponseWriter, req *http.Request)
 * @description:     This function handles /graph/history API request.
 * @related issues:  OWL-171
 * @param:           rw http.ResponseWriter
 * @param:           req *http.Request
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/12/2015
 * @last modified:   11/13/2015
 * @called by:       func configApiRoutes()
 */
func queryHistory(rw http.ResponseWriter, req *http.Request) {
	url := g.Config().Api.Query + "/graph/history"
	postByJSON(rw, req, url)
}

/**
 * @function name:   func getRequest(rw http.ResponseWriter, url string)
 * @description:     This function sends GET request to given URL.
 * @related issues:  OWL-159
 * @param:           rw http.ResponseWriter
 * @param:           url string
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/24/2015
 * @last modified:   11/24/2015
 * @called by:       func dashboardEndpoints(rw http.ResponseWriter, req *http.Request)
 *                    in query/http/api.go
 * @called by:       func dashboardEndpoints(rw http.ResponseWriter, req *http.Request)
 *                    in query/http/api.go
 */
func getRequest(rw http.ResponseWriter, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error =", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error =", err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Write(body)
}

/**
 * @function name:   func dashboardEndpoints(rw http.ResponseWriter, req *http.Request)
 * @description:     This function handles /api/endpoints API request.
 * @related issues:  OWL-159, OWL-171
 * @param:           rw http.ResponseWriter
 * @param:           req *http.Request
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/12/2015
 * @last modified:   11/24/2015
 * @called by:       func configApiRoutes()
 */
func dashboardEndpoints(rw http.ResponseWriter, req *http.Request) {
	url := g.Config().Api.Dashboard + req.URL.RequestURI()
	getRequest(rw, url)
}

/**
 * @function name:   func postByForm(rw http.ResponseWriter, req *http.Request, url string)
 * @description:     This function sends a POST request in Form format.
 * @related issues:  OWL-171
 * @param:           rw http.ResponseWriter
 * @param:           req *http.Request
 * @param:           url string
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/12/2015
 * @last modified:   11/13/2015
 * @called by:       func dashboardCounters(rw http.ResponseWriter, req *http.Request)
 *                   func dashboardChart(rw http.ResponseWriter, req *http.Request)
 */
func postByForm(rw http.ResponseWriter, req *http.Request, url string) {
	req.ParseForm()
	client := &http.Client{}
	resp, err := client.PostForm(url, req.PostForm)
	if err != nil {
		log.Println("Error =", err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Write(body)
}

/**
 * @function name:   func dashboardCounters(rw http.ResponseWriter, req *http.Request)
 * @description:     This function handles /api/counters API request.
 * @related issues:  OWL-171
 * @param:           rw http.ResponseWriter
 * @param:           req *http.Request
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/13/2015
 * @last modified:   11/13/2015
 * @called by:       func configApiRoutes()
 */
func dashboardCounters(rw http.ResponseWriter, req *http.Request) {
	url := g.Config().Api.Dashboard + "/api/counters"
	postByForm(rw, req, url)
}

/**
 * @function name:   func dashboardChart(rw http.ResponseWriter, req *http.Request)
 * @description:     This function handles /api/chart API request.
 * @related issues:  OWL-171
 * @param:           rw http.ResponseWriter
 * @param:           req *http.Request
 * @return:          void
 * @author:          Don Hsieh
 * @since:           11/13/2015
 * @last modified:   11/13/2015
 * @called by:       func configApiRoutes()
 */
func dashboardChart(rw http.ResponseWriter, req *http.Request) {
	url := g.Config().Api.Dashboard + "/chart"
	postByForm(rw, req, url)
}

func getAgentAliveData(hostnames []string, versions map[string]string, result map[string]interface{}) []cmodel.GraphLastResp {
	var queries []cmodel.GraphLastParam
	o := orm.NewOrm()
	var hosts []*Host
	_, err := o.Raw("SELECT hostname, agent_version FROM falcon_portal.host ORDER BY hostname ASC").QueryRows(&hosts)
	if err != nil {
		setError(err.Error(), result)
	} else {
		for _, host := range hosts {
			var query cmodel.GraphLastParam
			if !strings.Contains(host.Hostname, ".") && strings.Contains(host.Hostname, "-") {
				hostnames = append(hostnames, host.Hostname)
				versions[host.Hostname] = host.Agent_version
				query.Endpoint = host.Hostname
				query.Counter = "agent.alive"
				queries = append(queries, query)
			}
		}
	}
	s, err := json.Marshal(queries)
	if err != nil {
		setError(err.Error(), result)
	}
	url := g.Config().Api.Query + "/graph/last"
	reqPost, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(s)))
	if err != nil {
		setError(err.Error(), result)
	}
	reqPost.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqPost)
	if err != nil {
		setError(err.Error(), result)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	data := []cmodel.GraphLastResp{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		setError(err.Error(), result)
	}
	return data
}

func processAgentAliveData(data []cmodel.GraphLastResp, hostnames []string, versions map[string]string, result map[string]interface{}) {
	name := ""
	version := ""
	status := ""
	alive := 0
	countOfNormal := 0
	countOfWarn := 0
	countOfDead := 0
	anomalies := []interface{}{}
	items := []interface{}{}
	for key, row := range data {
		name = row.Endpoint
		var diff int64
		diff = 0
		var timestamp int64
		timestamp = 0
		status = "dead"
		alive = 0
		if name == "" {
			name = hostnames[key]
		} else {
			alive = int(row.Value.Value)
			timestamp = row.Value.Timestamp
			now := time.Now().Unix()
			diff = now - timestamp
		}
		version = versions[name]
		if alive > 0 {
			if diff > 3600 {
				status = "warm"
				countOfWarn++
			} else {
				status = "normal"
				countOfNormal++
			}
		} else {
			countOfDead++
		}
		item := map[string]interface{}{}
		item["id"] = strconv.Itoa(key + 1)
		item["hostname"] = name
		item["agent_version"] = version
		item["alive"] = alive
		item["timestamp"] = timestamp
		item["diff"] = diff
		item["status"] = status
		items = append(items, item)
		if diff > 60*60*24 && timestamp > 0 {
			anomalies = append(anomalies, item)
		}
	}
	var count = make(map[string]interface{})
	count["all"] = len(data)
	count["normal"] = countOfNormal
	count["warn"] = countOfWarn
	count["dead"] = countOfDead
	result["count"] = count
	result["items"] = items
}

func getAlive(rw http.ResponseWriter, req *http.Request) {
	var nodes = make(map[string]interface{})
	errors := []string{}
	var result = make(map[string]interface{})
	result["error"] = errors

	data := []cmodel.GraphLastResp{}
	hostnames := []string{}
	var versions = make(map[string]string)
	data = getAgentAliveData(hostnames, versions, result)
	processAgentAliveData(data, hostnames, versions, result)
	nodes["result"] = result
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	setResponse(rw, nodes)
}

func setStrategyTags(rw http.ResponseWriter, req *http.Request) {
	errors := []string{}
	var result = make(map[string]interface{})
	result["error"] = errors
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	json, err := simplejson.NewJson(buf.Bytes())
	if err != nil {
		setError(err.Error(), result)
	}

	var nodes = make(map[string]interface{})
	nodes, _ = json.Map()
	strategyID := ""
	tagName := ""
	tagValue := ""
	if value, ok := nodes["strategyID"]; ok {
		strategyID = value.(string)
		delete(nodes, "strategyID")
	}
	if value, ok := nodes["tagName"]; ok {
		tagName = value.(string)
		delete(nodes, "tagName")
	}
	if value, ok := nodes["tagValue"]; ok {
		tagValue = value.(string)
		delete(nodes, "tagValue")
	}

	if len(strategyID) > 0 && len(tagName) > 0 && len(tagValue) > 0 {
		strategyIDint, err := strconv.Atoi(strategyID)
		if err != nil {
			setError(err.Error(), result)
		} else {
			o := orm.NewOrm()
			var tag Tag
			sqlcmd := "SELECT * FROM falcon_portal.tags WHERE strategy_id=?"
			err = o.Raw(sqlcmd, strategyIDint).QueryRow(&tag)
			if err == orm.ErrNoRows {
				log.Println("tag not found")
				sql := "INSERT INTO tags(strategy_id, name, value, create_at) VALUES(?, ?, ?, ?)"
				res, err := o.Raw(sql, strategyIDint, tagName, tagValue, getNow()).Exec()
				if err != nil {
					setError(err.Error(), result)
				} else {
					num, _ := res.RowsAffected()
					log.Println("mysql row affected nums =", num)
					result["strategyID"] = strategyID
					result["action"] = "create"
				}
			} else if err != nil {
				setError(err.Error(), result)
			} else {
				log.Println("tag existed =", tag)
				sql := "UPDATE tags SET name = ?, value = ? WHERE strategy_id = ?"
				res, err := o.Raw(sql, tagName, tagValue, strategyIDint).Exec()
				if err != nil {
					setError(err.Error(), result)
				} else {
					num, _ := res.RowsAffected()
					log.Println("mysql row affected nums =", num)
					result["strategyID"] = strategyID
					result["action"] = "update"
				}
			}
		}
	} else {
		setError("Input value errors.", result)
	}
	nodes["result"] = result
	setResponse(rw, nodes)
}

func getTemplateStrategies(rw http.ResponseWriter, req *http.Request) {
	errors := []string{}
	var result = make(map[string]interface{})
	result["error"] = errors
	items := []interface{}{}
	countOfStrategies := 0
	arguments := strings.Split(req.URL.Path, "/")
	if arguments[len(arguments)-1] == "strategies" {
		templateID, err := strconv.Atoi(arguments[len(arguments)-2])
		if err != nil {
			setError(err.Error(), result)
		}
		o := orm.NewOrm()
		var strategyIDs []int64
		num, err := o.Raw("SELECT id FROM falcon_portal.strategy WHERE tpl_id = ? ORDER BY id ASC", templateID).QueryRows(&strategyIDs)
		if err != nil {
			setError(err.Error(), result)
		} else if num > 0 {
			countOfStrategies = int(num)
			var strategies = make(map[string]interface{})
			sids := ""
			for key, strategyID := range strategyIDs {
				sid := strconv.Itoa(int(strategyID))
				item := map[string]string{}
				item["templateID"] = strconv.Itoa(templateID)
				item["strategyID"] = sid
				strategies[sid] = item
				if key == 0 {
					sids = sid
				} else {
					sids += ", " + sid
				}
			}
			sqlcmd := "SELECT strategy_id, name, value FROM falcon_portal.tags WHERE strategy_id IN ("
			sqlcmd += sids
			sqlcmd += ") ORDER BY strategy_id ASC"
			var tags []*Tag
			_, err = o.Raw(sqlcmd).QueryRows(&tags)
			if err != nil {
				setError(err.Error(), result)
			} else {
				for _, tag := range tags {
					strategyID := strconv.Itoa(int(tag.StrategyId))
					strategy := strategies[strategyID].(map[string]string)
					strategy["tagName"] = tag.Name
					strategy["tagValue"] = tag.Value
					strategies[strategyID] = strategy
				}
			}
			for _, strategy := range strategies {
				items = append(items, strategy)
			}
		}
	}
	result["items"] = items
	result["count"] = countOfStrategies
	var nodes = make(map[string]interface{})
	nodes["result"] = result
	setResponse(rw, nodes)
}

func getPlatformJSON(nodes map[string]interface{}, result map[string]interface{}) {
	fcname := g.Config().Api.Name
	fctoken := getFctoken()
	url := g.Config().Api.Map + "/fcname/" + fcname + "/fctoken/" + fctoken
	url += "/show_active/yes/hostname/yes/pop_id/yes.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		setError(err.Error(), result)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		setError(err.Error(), result)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &nodes); err != nil {
		setError(err.Error(), result)
	}
}

func configAPIRoutes() {
	http.HandleFunc("/api/info", queryInfo)
	http.HandleFunc("/api/history", queryHistory)
	http.HandleFunc("/api/endpoints", dashboardEndpoints)
	http.HandleFunc("/api/counters", dashboardCounters)
	http.HandleFunc("/api/chart", dashboardChart)
	http.HandleFunc("/api/alive", getAlive)
	http.HandleFunc("/api/tags/update", setStrategyTags)
	http.HandleFunc("/api/templates/", getTemplateStrategies)
}
