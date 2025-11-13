package catalogue

import (
	"time"

	"github.com/NOLLYWOOD-COM/go-sdk/internal/httpclient"
)

type WorkSvc struct {
	httpClient httpclient.Client
}

type PeopleSvc struct {
	httpClient httpclient.Client
}

type ArticleSvc struct {
	httpClient httpclient.Client
}

type Person struct {
	ID             string     `json:"id"`
	HeadshotID     *string    `json:"headshotId"`
	Name           string     `json:"name"`
	Slug           string     `json:"slug"`
	Gender         *string    `json:"gender"`
	Age            *int       `json:"age"`
	Deceased       bool       `json:"deceased"`
	BirthName      *string    `json:"birthName"`
	BirthPlace     *string    `json:"birthPlace"`
	BirthDate      *FlexibleDate `json:"birthDate"`
	DeathDate      *FlexibleDate `json:"deathDate"`
	Aliases        []string   `json:"aliases"`
	Nationality    []string   `json:"nationality"`
	Bio            *string    `json:"bio"`
	HeightMetric   *int       `json:"heightMetric"`
	HeightImperial *string    `json:"heightImperial"`
	ExternalLinks  []struct {
		URL      string  `json:"url"`
		Label    *string `json:"label"`
		Icon     *string `json:"icon"`
		Platform string  `json:"platform"`
	} `json:"externalLinks"`
	Featured  bool      `json:"featured"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Work struct {
	ID              string     `json:"id"`
	ParentID        *string    `json:"parentId"`
	WorkType        string     `json:"workType"`
	Slug            string     `json:"slug"`
	Title           string     `json:"title"`
	OriginalTitle   string     `json:"originalTitle"`
	PosterID        *string    `json:"posterId"`
	BackdropID      *string    `json:"backdropId"`
	TrailerID       *string    `json:"trailerId"`
	VideoID         *string    `json:"videoId"`
	SpokenLanguages []string   `json:"spokenLanguages"`
	Languages       []string   `json:"languages"`
	SeasonCount     *int       `json:"seasonCount"`
	SeasonNumber    *int       `json:"seasonNumber"`
	EpisodeNumber   *int       `json:"episodeNumber"`
	Summary         *string    `json:"summary"`
	Synopsis        *string    `json:"synopsis"`
	ContentRating   *string    `json:"contentRating"`
	ReleaseDate     *string    `json:"releaseDate"`
	ReleaseYear     *int       `json:"releaseYear"`
	AirDate         *FlexibleDate  `json:"airDate"`
	StartDate       *FlexibleDate  `json:"startDate"`
	EndDate         *FlexibleDate  `json:"endDate"`
	Runtime         *int           `json:"runtime"`
	UserRating      *FlexibleFloat `json:"userRating"`
	CriticRating    *FlexibleFloat `json:"criticRating"`
	IsStreamable    bool       `json:"isStreamable"`
	IsInTheatre     bool       `json:"isInTheatre"`
	Budget          *float64   `json:"budget"`
	BudgetCurrency  *string    `json:"budgetCurrency"`
	Genres          []Genre    `json:"genres"`
	Featured        bool       `json:"featured"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt"`
}

type Genre struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	IsSubGenre  bool      `json:"isSubGenre"`
	Description *string   `json:"description"`
	ParentID    *string   `json:"parentId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Article struct {
	ID             string   `json:"id"`
	UserID         string   `json:"userId"`
	Title          string   `json:"title"`
	SeoTitle       *string  `json:"seoTitle"`
	SeoDescription *string  `json:"seoDescription"`
	SeoKeywords    []string `json:"seoKeywords"`
	Slug           string   `json:"slug"`
	CoverImageID   *string  `json:"coverImageId"`
	Summary        *string  `json:"summary"`
	Content        string   `json:"content"`
	Status         string   `json:"status"`
	Tags           []struct {
		ID         string    `json:"id"`
		ArticleID  string    `json:"articleId"`
		EntityType string    `json:"entityType"`
		EntityID   string    `json:"entityId"`
		CreatedAt  time.Time `json:"createdAt"`
	} `json:"tags"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	PublishedAt *time.Time `json:"publishedAt"`
	ArchivedAt  *time.Time `json:"archivedAt"`
}
