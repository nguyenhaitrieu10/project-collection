from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import os
import yaml

GITLAB_URL = "https://gitlab.example.com"
JENKINS_URL = "https://jenkins.example.com"

MAPPING = {
    "test": "sandbox",
    "qa": "qa"
}

def main():
    with open("input.yaml", "r") as f:
        data = yaml.load(f, yaml.FullLoader)

    envList = data["env"]
    projectList = data["projects"]
    dryRun = bool(data.get("dryRun", False))

    inputList = []
    for env in envList:
        for projectName in projectList:
            input = {}
            if not os.path.isdir("../../configs/%s" %projectName):
                print("Project Name not found")
                continue

            with open("../../configs/%s/config.yml" %projectName, "r") as f:
                data = yaml.load(f, yaml.FullLoader)

            input["env"] = env
            input["projectName"] = projectName
            input["gitUrl"] = GITLAB_URL + data["git"][7:-4].replace(":", "/") + "/hooks"
            input["jenkinsHook"] = JENKINS_URL + "/project/vn.%s.k8s/%s-pipeline" % (env, projectName)
            inputList.append(input)


    driver = webdriver.Chrome()
    driver.get(GITLAB_URL)
    eUser = driver.find_element_by_id("user_login")
    ePwd = driver.find_element_by_id("user_password")
    eUser.send_keys(os.getenv("GITLAB_USER", ""))
    ePwd.send_keys(os.getenv("GITLAB_PWD", ""))
    eLogin = driver.find_element_by_name("commit")
    eLogin.click()

    for input in inputList:
        print("%s  %s  %s  %s" %(input["projectName"], input["env"], input["gitUrl"], input["jenkinsHook"]))
        if dryRun:
            continue

        driver.get(input["gitUrl"])
        eHookUrl = driver.find_element_by_id("hook_url")
        eHookUrl.clear()
        eHookUrl.send_keys(input["jenkinsHook"])

        eBranch = driver.find_element_by_id("hook_push_events_branch_filter")
        eBranch.clear()
        eBranch.send_keys(MAPPING[input["env"]])

        eLogin = driver.find_element_by_name("commit")
        eLogin.click()

    driver.close()

if __name__ == "__main__":
    main()