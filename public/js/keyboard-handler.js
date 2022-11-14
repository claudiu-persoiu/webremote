const keyboardHandler = function (communication, buttonContainer, buttons) {

    const getActionFunction = function (value) {
        switch (value) {
            case "Shift":
                return toggleShift.toggle;
                break;
            case "Caps":
                return toggleCaps.toggle;
                break;
            case "Alt":
                return toggleAlt.toggle;
                break;
            case "Ctrl":
                return toggleCtrl.toggle;
                break;
            default:
                return buttonAction;
        }
    };
    const updateKeyboard = function (key) {
        key = key || "default";
        const buttons = buttonContainer.querySelectorAll("button");
        buttons.forEach(function (button) {
            button.innerHTML = button.dataset[key] || button.dataset.default;
        });
    };
    const toggle = function (key) {
        let flag = false;
        return {
            toggle: function () {
                feedback();
                flag = !flag;
                document.querySelectorAll('.class-for-' + key + ' button').forEach(function (button) {
                    button.className = flag && "selected";
                });
                updateKeyboard(flag && key);
            },
            value: function () {
                return flag;
            }
        };
    };

    const toggleShift = toggle("shift");
    const toggleCaps = toggle("caps");
    const toggleAlt = toggle("alt");
    const toggleCtrl = toggle("ctrl");

    const buttonAction = function () {
        feedback();
        const value = this.value,
            command = [value];

        if (toggleShift.value()) {
            toggleShift.toggle();
            command.unshift("Shift");
        } else if (toggleCaps.value() && this.value != this.dataset.caps) {
            command.unshift("Shift");
        }

        if (toggleAlt.value()) {
            command.unshift("Alt");
        }

        if (toggleCtrl.value()) {
            command.unshift("Ctrl");
        }

        communication.send({
            type: "keyboard",
            commands: command
        });
    };

    const drawKeyboard = function () {
        buttonContainer.innerHTML = "";
        let row, key, defaultValue, button, actionFunction;
        for (const rowIndex in buttons) {
            row = document.createElement("div");
            row.className = "row-container row-" + rowIndex;
            for (const keyIndex in buttons[rowIndex]) {
                key = buttons[rowIndex][keyIndex];
                defaultValue = key.default;
                button = document.createElement("button");
                button.innerHTML = defaultValue;
                button.value = defaultValue;
                button.onclick = getActionFunction(defaultValue);
                button.dataset.default = defaultValue;
                button.dataset.shift = key.shift || defaultValue;
                button.dataset.caps = key.caps || defaultValue;
                const div2 = document.createElement("div");
                div2.appendChild(button);
                div2.className = "child-container class-for-" + defaultValue.toLowerCase();
                row.appendChild(div2);
            }
            buttonContainer.appendChild(row);
        }
    };

    drawKeyboard();
};
