*** Variables ***

*** Keywords ***
# Open Browser for OS
#     ${SYSTEM}=                   Evaluate                     platform.system()                                    platform
#     Run Keyword If               '${SYSTEM}' == 'Windows'     Open Chrome With Options
#     ...                          ELSE                         Open Mac Chrome
#     Set Selenium Speed           0

Open Chrome
    Open Browser                 http://localhost      Chrome
    Maximize Browser Window

# Open Chrome With Options
#     ${OPTIONS}=                  Evaluate                     sys.modules['selenium.webdriver'].ChromeOptions()    sys, selenium.webdriver
#     Call Method                  ${OPTIONS}                   add_experimental_option                              useAutomationExtension     ${FALSE}
#     Call Method                  ${OPTIONS}                   add_argument                                         start-maximized
#     Create WebDriver             Chrome                       chrome_options=${OPTIONS}
