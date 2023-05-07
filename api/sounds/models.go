package sounds

import "time"

type SoundsResponse struct {
	Total    int `json:"total"`
	Size     int `json:"size"`
	Page     int `json:"page"`
	LastPage int `json:"lastPage"`
	Items    []struct {
		ID    string `json:"id"`
		Owner struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
			Verified bool   `json:"verified"`
		} `json:"owner"`
		Title       string    `json:"title"`
		Category    string    `json:"category"`
		Dmca        bool      `json:"dmca"`
		Tags        []string  `json:"tags"`
		ImagePath   string    `json:"imagePath"`
		ThumbnailS  string    `json:"thumbnailS"`
		ThumbnailM  string    `json:"thumbnailM"`
		Path        string    `json:"path"`
		OggPath     string    `json:"oggPath"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		Anonymous   bool      `json:"anonymous"`
		Social      struct {
			LikeCount   int `json:"likeCount"`
			SharedCount int `json:"sharedCount"`
		} `json:"social"`
		Properties struct {
			Duration         int     `json:"duration"`
			Loudness         float64 `json:"loudness"`
			MaximumAmplitude float64 `json:"maximumAmplitude"`
		} `json:"properties"`
		Permalink        string    `json:"permalink"`
		Explicit         bool      `json:"explicit"`
		UpdatedAt        time.Time `json:"updatedAt"`
		ModerationStatus string    `json:"moderationStatus"`
	} `json:"items"`
	Tags []struct {
		Name       string `json:"name"`
		SoundCount int    `json:"soundCount"`
	} `json:"tags"`
	Dmca []struct {
		Name       bool `json:"name"`
		SoundCount int  `json:"soundCount"`
	} `json:"dmca"`
	Permalink string `json:"permalink"`
}
