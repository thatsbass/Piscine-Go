#!/bin/bash

VAR_ENV_INTERVIEW_KEY=$(grep "INTERVIEW" streets/Buckingham_Place | sed 's/SEE INTERVIEW #//')

echo $VAR_ENV_INTERVIEW_KEY

var="interviews/interview-"

cat "$var$VAR_ENV_INTERVIEW_KEY"

echo $MAIN_SUSPECT
