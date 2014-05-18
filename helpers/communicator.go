package lifecycle_helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abhishekkr/gol/golerror"
	"github.com/abhishekkr/gol/golzmq"
)

func ZmqRead(key_type string, key string) string {
	rep_port, err_rep_port := strconv.Atoi(*LifeCycleConfig["db_rep_port"])
	req_port, err_req_port := strconv.Atoi(*LifeCycleConfig["db_req_port"])
	if err_rep_port != nil || err_req_port != nil {
		golerror.Boohoo("Port parameters to bind, error-ed while conversion to number.", true)
	}

	sock := golzmq.ZmqRequestSocket(*LifeCycleConfig["db_uri"], rep_port, req_port)
	response, _ := golzmq.ZmqRequest(sock, "read", key_type, key)
	return response
}

func ChildNodes(parent_node string) []string {
	key_parent_node := fmt.Sprintf("key::%s", parent_node)
	key_nodes := ZmqRead("default", key_parent_node)
	nodes := strings.Split(key_nodes, ",")
	for idx, node := range nodes {
		node_fields := strings.Split(node, fmt.Sprintf("%s:", key_parent_node))
		if len(node_fields) > 1 {
			nodes[idx] = strings.Join(node_fields[1:], "")
		}
	}
	return nodes
}

/*
	abkzeromq.ZmqReq(*req_port, *rep_port, "read", "default", "key::node")
	abkzeromq.ZmqReq(*req_port, *rep_port, "read", "default", "key::node:h4ck3r.edu")
	abkzeromq.ZmqReq(*req_port, *rep_port, "read", "tsds", "node:h4ck3r.edu")
*/
