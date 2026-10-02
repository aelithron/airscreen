# airscreen
a little utility that sends hack club ysws submissions (from an airtable base) to [robert](https://telescreen.hackclub.com/robert) in telescreen for fraud checking! \
written in golang and fully compliant with the robert api. **exclusively for software!**
## using
you need a robert api key for the program you run, as well as an airtable pat, base id, and table id. \
said table's fields must include the following:
| field name | description | | required |
| --- | --- | --- |
| Code URL | link to the git repo | yes |
| Playable URL | link to the deployed demo | yes |
| Hackatime Projects | comma-seperated list of hackatime projects | yes |
| Project Name | the project's name | no (will use the first hackatime project's name if omitted) |
| Slack ID | the submitter's slack id | no if hackatime id |
| Hackatime ID | the submitter's numerical hackatime id | no if slack id |