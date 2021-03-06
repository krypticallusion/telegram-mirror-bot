package aria

import "github.com/zyxar/argo/rpc"

func Download(url string) (string, error) {
	return client.AddURI([]string{url})
}

func GetStatus(gid string) (si rpc.StatusInfo) {
	si, _ = client.TellStatus(gid)
	return
}

func Cancel(gid string) error {
	_, err := client.Remove(gid)

	if err != nil {
		return err
	}

	return nil
}
