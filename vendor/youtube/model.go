package youtube

type YoutubeVideo struct {
	UploadDate    string      `json:"upload_date"`
	Protocol      string      `json:"protocol"`
	Creator       interface{} `json:"creator"`
	Series        interface{} `json:"series"`
	FormatNote    string      `json:"format_note"`
	Chapters      interface{} `json:"chapters"`
	SeasonNumber  interface{} `json:"season_number"`
	LikeCount     int         `json:"like_count"`
	Duration      int         `json:"duration"`
	Fulltitle     string      `json:"fulltitle"`
	PlayerURL     string      `json:"player_url"`
	ID            string      `json:"id"`
	ViewCount     int         `json:"view_count"`
	Playlist      string      `json:"playlist"`
	Title         string      `json:"title"`
	Filename      string      `json:"_filename"`
	Format        string      `json:"format"`
	Ext           string      `json:"ext"`
	PlaylistIndex int         `json:"playlist_index"`
	DislikeCount  int         `json:"dislike_count"`
	AverageRating float64     `json:"average_rating"`
	Abr           int         `json:"abr"`
	UploaderURL   string      `json:"uploader_url"`
	Subtitles     struct {
	} `json:"subtitles"`
	AgeLimit           int         `json:"age_limit"`
	Annotations        interface{} `json:"annotations"`
	WebpageURLBasename string      `json:"webpage_url_basename"`
	Filesize           int         `json:"filesize"`
	DisplayID          string      `json:"display_id"`
	AutomaticCaptions  struct {
	} `json:"automatic_captions"`
	Description        string      `json:"description"`
	Tags               []string    `json:"tags"`
	RequestedSubtitles interface{} `json:"requested_subtitles"`
	StartTime          interface{} `json:"start_time"`
	Tbr                float64     `json:"tbr"`
	PlaylistID         string      `json:"playlist_id"`
	Uploader           string      `json:"uploader"`
	FormatID           string      `json:"format_id"`
	EpisodeNumber      interface{} `json:"episode_number"`
	UploaderID         string      `json:"uploader_id"`
	Categories         []string    `json:"categories"`
	PlaylistTitle      interface{} `json:"playlist_title"`
	Thumbnails         []struct {
		URL string `json:"url"`
		ID  string `json:"id"`
	} `json:"thumbnails"`
	License      string      `json:"license"`
	AltTitle     interface{} `json:"alt_title"`
	URL          string      `json:"url"`
	ExtractorKey string      `json:"extractor_key"`
	Vcodec       string      `json:"vcodec"`
	HTTPHeaders  struct {
		AcceptCharset  string `json:"Accept-Charset"`
		AcceptLanguage string `json:"Accept-Language"`
		AcceptEncoding string `json:"Accept-Encoding"`
		Accept         string `json:"Accept"`
		UserAgent      string `json:"User-Agent"`
	} `json:"http_headers"`
	Thumbnail  string      `json:"thumbnail"`
	IsLive     interface{} `json:"is_live"`
	Extractor  string      `json:"extractor"`
	EndTime    interface{} `json:"end_time"`
	WebpageURL string      `json:"webpage_url"`
	Formats    []struct {
		Asr         int         `json:"asr,omitempty"`
		Tbr         float64     `json:"tbr,omitempty"`
		Protocol    string      `json:"protocol"`
		Format      string      `json:"format"`
		FormatNote  string      `json:"format_note"`
		Height      interface{} `json:"height,omitempty"`
		ManifestURL string      `json:"manifest_url,omitempty"`
		FormatID    string      `json:"format_id"`
		Container   string      `json:"container,omitempty"`
		Language    interface{} `json:"language,omitempty"`
		HTTPHeaders struct {
			AcceptCharset  string `json:"Accept-Charset"`
			AcceptLanguage string `json:"Accept-Language"`
			AcceptEncoding string `json:"Accept-Encoding"`
			Accept         string `json:"Accept"`
			UserAgent      string `json:"User-Agent"`
		} `json:"http_headers"`
		URL        string      `json:"url"`
		Vcodec     string      `json:"vcodec"`
		Abr        int         `json:"abr,omitempty"`
		Width      interface{} `json:"width,omitempty"`
		Ext        string      `json:"ext"`
		Filesize   int         `json:"filesize,omitempty"`
		Fps        interface{} `json:"fps,omitempty"`
		Acodec     string      `json:"acodec"`
		PlayerURL  string      `json:"player_url,omitempty"`
		Resolution string      `json:"resolution,omitempty"`
	} `json:"formats"`
	Acodec   string `json:"acodec"`
	NEntries int    `json:"n_entries"`
}
