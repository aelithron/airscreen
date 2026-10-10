# airscreen ![IMG](https://hackatime-badge.hackclub.com/U08RJ1PEM7X/airscreen)

<a href="https://notbyai.fyi" target="_blank">
  <img src="./not-by-ai.svg" alt="Developed by a human, not by AI!">
</a>

a little utility that sends hack club ysws submissions (from an airtable base) to robert (in telescreen) for fraud checking! \
written in golang and fully compliant with the robert api. **exclusively for software!** \
note: this will also only send projects that have been approved in "first-pass" review, but not submitted to unified. more modes may be coming eventually, idk!
## using
you need a robert api key for the program you run, as well as an airtable pat (with base reading permissions), base id, and table id. the cli will ask you for these on first run, you can just paste them in! alternatively, you can fill them into `~/.airscreen/config.json` in the format of the [example config](https://github.com/aelithron/airscreen/blob/main/example.config.json).

the airtable's fields must include the following, named exactly as specified:

| field name | description | required |
| --- | --- | --- |
| Code URL | link to the git repo | yes |
| Playable URL | link to the deployed demo | yes |
| Hackatime Project Name | comma-seperated list of hackatime projects, or just a name if only one | yes |
| Automation - First Submitted At | empty for projects that haven't been sent to unified, or a date/time for submitted ones | yes |
| Review Status | "Approved" for projects that have passed "first-pass" review, anything else for non-approved | yes |
| Project Name | the project's name | no (will use the first hackatime project's name if omitted) |
| Slack ID | the submitter's slack id | no if hackatime id |
| Hackatime ID | the submitter's numerical hackatime id | no if slack id |

then, simply run the [airscreen binary](https://github.com/aelithron/airscreen/releases/latest) for your version, you may be able to just double-click it on some operating systems (some may require you run it from the terminal, just `./<name>` in the binary's directory to do so). 