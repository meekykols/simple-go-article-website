document.addEventListener('DOMContentLoaded', function () {
    var buttons = document.querySelectorAll('.toggle-btn');

    buttons.forEach(function (button) {
        button.addEventListener('click', function () {
            var content = button.nextElementSibling;
            content.classList.toggle('hidden');

            // Toggle text color between blue and original color
            var article = content.parentElement;
            article.classList.toggle('blue-text');
        });
    });
});

var currentYear = new Date().getFullYear();
document.getElementById("copyrightYear").innerHTML=currentYear;