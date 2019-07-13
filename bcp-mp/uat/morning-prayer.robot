*** Settings ***
Library                 SeleniumLibrary
#Library                lib/Util.py
Resource                base.robot
Suite Setup             Run Keywords    Open Chrome
# ...                     AND             Go To   http://localhost:3000
# ...                     AND             User Logs In To Global Access   ${user_name}    ${user_password}
Suite Teardown        Close Browser

*** Variables ***
# ${USER_NAME}        iblbeneficiary@gmail.com
# ${USER_PASSWORD}    FlynnMonkeyHat4

*** Test Cases ***

User Can See Page
    Given User Opens Page
    Then User Can See Page Title

*** Keywords ***

User Opens Page
    Go To   http://localhost:3000

User Can See Page Title
    Title Should Be  Morning Prayer  
