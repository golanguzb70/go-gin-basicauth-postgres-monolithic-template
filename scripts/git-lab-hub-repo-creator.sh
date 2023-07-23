#!/bin/bash

# Function to install jq if not already installed
install_jq() {
    if command -v jq &>/dev/null; then
        echo "jq is already installed."
    else
        echo "Installing jq..."
        if [[ "$(uname)" == "Darwin" ]]; then
            # macOS
            brew install jq
        elif [[ "$(uname)" == "Linux" ]]; then
            # Linux (Debian/Ubuntu)
            sudo apt-get update
            sudo apt-get install -y jq
        else
            echo "Error: Unsupported operating system. Please install jq manually."
            exit 1
        fi
    fi
}

# Function to create a new repository on GitHub using GitHub API
create_github_repository() {
    # GitHub API endpoint
    GITHUB_API="https://api.github.com/user/repos"

    # Prompt for GitHub username and personal access token if not provided
    if [ "$GITHUB_USERNAME" = "" ]; then
        echo -n "Enter your GitHub username: "
        read -r GITHUB_USERNAME
    fi

    if [ "$GITHUB_PERSONAL_ACCESS_TOKEN" = "" ]; then
        echo -n "Enter your GitHub personal access token: "
        read -r GITHUB_PERSONAL_ACCESS_TOKEN
    fi

    # GitHub repository name
    echo -n "Enter the GitHub repository name: "
    read -r REPO_NAME

    # GitHub repository description (Optional)
    echo -n "Enter the GitHub repository description (leave blank for no description): "
    read -r REPO_DESCRIPTION

    # GitHub repository visibility (public or private)
    VISIBILITIES="private public"
    VISIBILITY=""
    select vis in $VISIBILITIES 
    do 
        if [ "$vis" == "private" ]; then
            VISIBILITY="private"
        else 
            VISIBILITY="public"
        fi
        break
    done

    # Create the GitHub repository using GitHub API
    response=$(curl --header "Authorization: token $GITHUB_PERSONAL_ACCESS_TOKEN" --request POST \
        --data "{\"name\":\"$REPO_NAME\",\"description\":\"$REPO_DESCRIPTION\",\"private\":$([ "$VISIBILITY" == "private" ] && echo true || echo false)}" \
        "$GITHUB_API" | tr -d '\n')

    # Check the response for errors and print appropriate messages
    if [ $? -eq 0 ]; then
        if [ "$(echo "$response" | jq '.message')" = "null" ]; then
            echo "Repository created successfully."
        else
            echo "Error: $(echo "$response" | jq -r '.message')"
            exit 1
        fi
    else
        echo "Error: Failed to communicate with GitHub API."
        exit 1
    fi
}

# Function to create a new repository on GitLab using GitLab API
create_gitlab_repository() {
    # GitLab API endpoint
    GITLAB_API="https://gitlab.com/api/v4/projects"

    # Prompt for GitLab username and personal access token if not provided
    if [ "$GITLAB_USERNAME" = "" ]; then
        echo -n "Enter your GitLab username: "
        read -r GITLAB_USERNAME
    fi

    if [ "$GITLAB_PERSONAL_ACCESS_TOKEN" = "" ]; then
        echo -n "Enter your GitLab personal access token: "
        read -r GITLAB_PERSONAL_ACCESS_TOKEN
    fi

    # GitLab repository name
    echo -n "Enter the GitLab repository name: "
    read -r REPO_NAME

    # GitLab repository description (Optional)
    echo -n "Enter the GitLab repository description (leave blank for no description): "
    read -r REPO_DESCRIPTION

    # GitLab repository visibility (public, internal, or private)
    VISIBILITIES="private public"
    VISIBILITY=""
    select vis in $VISIBILITIES 
    do 
        if [ "$vis" == "private" ]; then
            VISIBILITY="private"
        else 
            VISIBILITY="public"
        fi
        break
    done


    # Create the GitLab repository using GitLab API
    local response
    response=$(curl --header "PRIVATE-TOKEN: $GITLAB_PERSONAL_ACCESS_TOKEN" --request POST \
        --data "name=$REPO_NAME&description=$REPO_DESCRIPTION&visibility=$VISIBILITY" \
        "$GITLAB_API")

    if [ $? -eq 0 ]; then
        if [ "$(echo "$response" | jq '.message')" = "null" ]; then
            echo "Repository created successfully on GitLab."
        else
            echo "Error: $(echo "$response" | jq -r '.message')"
            exit 1
        fi
    else
        echo "Error: Failed to communicate with GitLab API."
        exit 1
    fi
}

# Function to connect the current folder to the remote GitHub repository
connect_to_github_repository() {
    git init
    git add -A
    git commit -m "initial commit"
    # Open main branch
    git branch -M main
    # Add the remote repository URL
    git remote add origin "git@github.com:$GITHUB_USERNAME/$REPO_NAME.git"
    # Push the local repository to the remote repository
    git push -u origin main

    echo "Connected the current folder to the remote repository successfully."
}

# Function to connect the current folder to the remote GitLab repository
connect_to_gitlab_repository() {
    git init --initial-branch=main
    git remote add origin git@gitlab.com:$GITLAB_USERNAME/$REPO_NAME.git
    git add .
    git commit -m "Initial commit"
    git push --set-upstream origin main

    echo "Connected the current folder to the remote GitLab repository successfully."
}

# Function to determine the choice of the user (GitHub or GitLab)
choose_repository_type() {
    echo "Choose the type of repository:"
    options="GitHub GitLab Cancel"
    select opt in $options
    do
        if [ "$opt" == "GitHub" ]; then
            create_github_repository
            connect_to_github_repository
            break
        elif [ "$opt" == "GitLab" ]; then
            create_gitlab_repository
            connect_to_gitlab_repository
            break
        elif [ "$opt" == "Cancel" ]; then
            echo "Operation canceled."
            exit
        else
            echo "Invalid option. Please choose again."
        fi
    done
}

# Call the function to install jq if needed
install_jq

# Call the function to choose the type of repository and perform the necessary steps
choose_repository_type
