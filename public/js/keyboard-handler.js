var keyboardHandler = function (communication, buttonContainer, buttons) {

      var toggle = function (key) {
          var flag = false;
          return {
              toggle: function () {
                  flag = !flag;
                  document.querySelectorAll('.class-for-'+key+' button').forEach(function (button) {
                    button.className = flag && "selected";
                  });
                  updateKeyboard(flag && key);
              },
              value: function () {
                  return flag;
              }
          };
      }

      var toggleShift = toggle("shift");
      var toggleCapse = toggle("caps");
      var toggleAlt = toggle("alt");
      var toggleCtrl = toggle("ctrl");

      var buttonAction = function () {
          var value = this.value,
              command = [value];

          if (toggleShift.value()) {
            toggleShift.toggle();
            command.unshift("Shift");
          } else if (toggleCapse.value() && this.value != this.dataset.caps) {
            command.unshift("Shift");
          }

          if (toggleAlt.value()) {
            command.unshift("Alt");
          }

          if (toggleCtrl.value()) {
            command.unshift("Ctrl");
          }

          console.log(command);
          communication.send({
            type: "keyboard",
            commands: command
          });
      };

      var updateKeyboard = function (key) {
          key = key || "default";
          var buttons = buttonContainer.querySelectorAll("button");
          buttons.forEach(function (button) {
            button.innerHTML = button.dataset[key] || button.dataset.default;
          });
      };

      var drawKeyboard = function () {
          buttonContainer.innerHTML = "";
          var row, key, defaultValue, button, actionFunction;
          for (var rowIndex in buttons) {
              row = document.createElement("div");
              row.className = "row-container row-" + rowIndex;
              for (var keyIndex in buttons[rowIndex]) {
                  key = buttons[rowIndex][keyIndex];
                  defaultValue = key.default;
                  button = document.createElement("button");
                  button.innerHTML = defaultValue;
                  button.value = defaultValue;
                  button.onclick = getActionFunction(defaultValue);
                  button.dataset.default = defaultValue;
                  button.dataset.shift = key.shift || defaultValue;
                  button.dataset.caps = key.caps || defaultValue;
                  var div2 = document.createElement("div");
                  div2.appendChild(button);
                  div2.className = "child-container class-for-" + defaultValue.toLowerCase();
                  row.appendChild(div2);
              }
              buttonContainer.appendChild(row);
          }
      };

      var getActionFunction = function (value) {
          switch (value) {
            case "Shift":
              return toggleShift.toggle;
              break;
            case "Caps":
              return toggleCapse.toggle;
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

      drawKeyboard();
};
