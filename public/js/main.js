// Purpose: This file contains the JavaScript code for the single-page application (SPA) loader.
//
// This code is responsible for loading content from the server and updating the page without a full page reload.

spaLoader()

function spaLoader() {
    document.addEventListener("DOMContentLoaded", () => {
        attachAllListeners()

        window.addEventListener("popstate", () => {
            fetchContentJSON(window.location.pathname)
        }) // Listen for back/forward button clicks and fetch content
    })
}


function fetchContentJSON(url) {
    const contentDiv = document.getElementById("content")
    fetch(url)
        .then((response) => response.json())  // Change this to parse JSON instead of text
        .then((data) => {
            // Parse the HTML content
            const parser = new DOMParser()
            const doc = parser.parseFromString(data.html, "text/html")

            // Update the content
            contentDiv.innerHTML = doc.documentElement.innerHTML

            // Update the head
            document.title = data.title

            // Update the description meta tag
            const currentDescription = document.querySelector('meta[name="description"]')
            if (currentDescription) {
                currentDescription.setAttribute("content", data.description)
            } else {
                // If the description meta tag doesn't exist, create it
                const metaDescription = document.createElement('meta')
                metaDescription.name = "description"
                metaDescription.content = data.description
                document.head.appendChild(metaDescription)
            }

            attachAllListeners() // Reattach listeners to new content
        })
        .catch((error) => {
            console.error('Error fetching content:', error)
            contentDiv.innerHTML = '<p>Error loading content. Please try again.</p>'
        })
}


function attachButtonListener(button) {
    button.addEventListener(
        "click",
        (e) => {
            e.preventDefault() // Prevent default button behavior
            e.stopPropagation() // Stop event from bubbling up
            const url = button.getAttribute("data-href")
            if (url && url.startsWith("/")) {
                history.pushState(null, "", url)
                fetchContent(url)
            }

            return false // Prevent onclick from firing
        },
        true,
    ) // Use capturing to ensure this runs before onclick
}

function attachLinkListener(link) {
    link.addEventListener(
        "click",
        (e) => {
            e.preventDefault() // Prevent default link behavior
            e.stopPropagation() // Stop event from bubbling up
            const url = link.getAttribute("href")
            if (url && url.startsWith("/")) {
                history.pushState(null, "", url)
                fetchContentJSON("/nav" + url)
            }
            return false // Prevent onclick from firing
        },
        true,
    ) // Use capturing to ensure this runs before onclick
}


function attachAllListeners() {
    document
        .querySelectorAll('a[href^="/"]')
        .forEach(attachLinkListener)
}

