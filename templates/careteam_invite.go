package templates

import "./../models"

const _CareteamInviteSubjectTemplate string = `Diabetes care team invitation`
const _CareteamInviteBodyTemplate string = `
<html>
  <head>
    <meta name='viewport' content='width=device-width'/>
    <meta http-equiv='Content-Name' content='text/html; charset=UTF-8'/>
    <title>Diabetes care team invitation</title>
    <link href='http://fonts.googleapis.com/css?family=Nunito:400,300,700' rel='stylesheet' type='text/css'>
  </head>

  <body style='background-color: #FFFFFF'>

    <div class='container' style='background-color:#F5F5F5; padding:20px; margin:0 auto; max-width:500px'>
      <div align='center' style='padding:10px; margin:0;'>
        <a href='{{ .BlipURL }}/'><img width='75' height='75' src='http://drive.google.com/uc?export=view&id=0BwI0YrjnbmtXTUoxc1JRdDViaWc' /></a>
      </div>

      <div align='center'>
        <p style='font-family: Nunito, sans-serif, Helvetica Neue, Helvetica; font-weight:300; font-size: 14px; color:#000; line-height:1.1; padding:25px 0 15px; margin:0;'>Hey there!</p>
        <p style='font-family: Nunito, sans-serif, Helvetica Neue, Helvetica; font-weight:300; font-size: 14px; color:#000; line-height:1.1; padding:0 0 15px; margin:0;'>{{ .CareteamName }} invited you to be on their care team. Please click the link below to accept and see {{ .CareteamName }}’s data.</p>
      </div>

      <br>

      {{if .IsExistingUser}}
      <div align='center' style='padding:0;'>
        <a style='background-color:#627CFB; font-family: Nunito, sans-serif, Helvetica Neue, Helvetica; font-weight:400; font-size: 14px; color:#FFFFFF; padding:10px 20px; margin:0; border-radius:20px; text-decoration: none;' href='{{ .BlipURL }}/#/login?inviteEmail={{ .Email }}&inviteKey={{ .Key }}'>Join {{ .CareteamName }}'s Care Team</a>
      </div>
      {{else}}
      <div align='center' style='padding:0;'>
        <a style='background-color:#627CFB; font-family: Nunito, sans-serif, Helvetica Neue, Helvetica; font-weight:400; font-size: 14px; color:#FFFFFF; padding:10px 20px; margin:0; border-radius:20px; text-decoration: none;' href='{{ .BlipURL }}/#/signup?inviteEmail={{ .Email }}&inviteKey={{ .Key }}'>Join {{ .CareteamName }}'s Care Team</a>
      </div>
      {{end}}

      <br>

      <div align='center'>
        <p style='font-family: Nunito, sans-serif, Helvetica Neue, Helvetica; font-weight:300; font-size: 14px; color:#000; line-height:1.1; padding:15px 0 40px; margin:0;'>Sincerely,<br>The Tidepool Team</p>
      </div>

      <div align='center' style='padding:10px 10px 0; margin:0;'>
        <a href='https://www.tidepool.org'><img width='139.01162791' height='15' src='http://drive.google.com/uc?export=view&id=0BwI0YrjnbmtXYkdQS0xqaThyTGc'/></a>
      </div>
      <div align='center' style='font-family: Nunito, sans-serif, Helvetica Neue, Helvetica; font-weight:300; font-size: 12px; color:#444; line-height:1.8; padding:5px 0 0 0; margin:0;'>
        <a style='margin:0; display:block; text-decoration:none; color:#444' href='https://www.tidepool.org'>tidepool.org</a>
      </div>
    </div>

  </body>
</html>
`

func NewCareteamInviteTemplate() (models.Template, error) {
	return models.NewPrecompiledTemplate(models.TemplateNameCareteamInvite, _CareteamInviteSubjectTemplate, _CareteamInviteBodyTemplate)
}