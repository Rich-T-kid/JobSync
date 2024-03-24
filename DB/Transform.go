package DB
import( "time")





func parseTimeString(timeStr string) (time.Time, error) {
    // Adjust the layout according to the timestamp format
    layout := "2006-01-02 15:04:05"
    return time.Parse(layout, timeStr)
}
