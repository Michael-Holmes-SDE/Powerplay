package schedule

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/csv"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/parser"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type RefScheduleRow struct {
	Start           string `csv:"Start Date and Time"`
	DurationHours   string `csv:"Duration Hours"`
	DurationMinutes string `csv:"Duration Minutes"`
	Location        string `csv:"Arena/Rink"`
	Level           string `csv:"Game Level"`
	Home            string `csv:"Home Team"`
	Away            string `csv:"Away Team"`
}

func handleRef(c *fiber.Ctx) error {
	games, seasonConfig := parser.ReadGames("spring_2024")

	refSchedule := make([]RefScheduleRow, 0)

	for i, game := range games {
		for _, league := range seasonConfig.Leagues {
			for _, team := range league.Teams {
				if game.Team1Name == team.Name {
					games[i].League = league.Name
					break
				}
			}
		}

		row := RefScheduleRow{
			Start:           game.Start.Format("1/2/06 3:04 PM"),
			DurationHours:   "1",
			DurationMinutes: "15",
			Location:        "George S. Eccles Ice Center",
			Level:           fmt.Sprintf("%s League", game.League),
			Home:            game.Team1Name,
			Away:            game.Team2Name,
		}

		refSchedule = append(refSchedule, row)
	}

	csv.GenerateCsv(refSchedule, "ref_schedule.csv")
	return responder.NotYetImplemented(c)
}
