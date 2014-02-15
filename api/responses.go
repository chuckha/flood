package api

type Droplet struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	ImageId          int    `json:"image_id"`
	SizeId           int    `json:"size_id"`
	RegionId         int    `json:"region_id"`
	BackupsActive    bool   `json:"backups_active"`
	IpAddress        string `json:"ip_address"`
	PrivateIpAddress string `json:"private_ip_address"`
	Locked           bool   `json:"locked"`
	Status           string `json:"status"`
	CreatedAt        string `json:"created_at"`
}

type DropletListResponse struct {
	Status   string    `json:"status"`
	Droplets []Droplet `json:"droplets"`
}

type Size struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SizeListResponse struct {
	Status string `json:"status`
	Sizes  []Size `json:"sizes"`
}

type Image struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
}

type ImageListResponse struct {
	Status string  `json:"status`
	Sizes  []Image `json:"images"`
}
