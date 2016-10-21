var keyboardHandler = function (communication, buttonContainer) {

      var buttons = [
          [
            {
              "default": "Esc"
            },
            {
              "default": "F1"
            },
            {
              "default": "F2"
            },
            {
              "default": "F3"
            },
            {
              "default": "F4"
            },
            {
              "default": "F5"
            },
            {
              "default": "F6"
            },
            {
              "default": "F7"
            },
            {
              "default": "F8"
            },
            {
              "default": "F9"
            },
            {
              "default": "F10"
            },
            {
              "default": "F11"
            },
            {
              "default": "F12"
            }
          ],
          [
            {
              "default": "`",
              "shift": "~"
            },
            {
              "default": "1",
              "shift": "!"
            },
            {
              "default": "2",
              "shift": "@"
            },
            {
              "default": "3",
              "shift": "#"
            },
            {
              "default": "4",
              "shift": "$"
            },
            {
              "default": "5",
              "shift": "%"
            },
            {
              "default": "6",
              "shift": "^"
            },
            {
              "default": "7",
              "shift": "&"
            },
            {
              "default": "8",
              "shift": "*"
            },
            {
              "default": "9",
              "shift": "("
            },
            {
              "default": "0",
              "shift": ")"
            },
            {
              "default": "-",
              "shift": "_"
            },
            {
              "default": "=",
              "shift": "+"
            },
            {
              "default": "Bksp"
            }
          ],
          [
            {
              "default": "Tab"
            },
            {
              "default": "q",
              "shift": "Q",
              "caps": "Q"
            },
            {
              "default": "w",
              "shift": "W",
              "caps": "W"
            },
            {
              "default": "e",
              "shift": "E",
              "caps": "E"
            },
            {
              "default": "r",
              "shift": "R",
              "caps": "R"
            },
            {
              "default": "t",
              "shift": "T",
              "caps": "T"
            },
            {
              "default": "y",
              "shift": "Y",
              "caps": "Y"
            },
            {
              "default": "u",
              "shift": "U",
              "caps": "U"
            },
            {
              "default": "i",
              "shift": "I",
              "caps": "I"
            },
            {
              "default": "o",
              "shift": "O",
              "caps": "O"
            },
            {
              "default": "p",
              "shift": "P",
              "caps": "P"
            },
            {
              "default": "[",
              "shift": "{"
            },
            {
              "default": "]",
              "shift": "}"
            },
            {
              "default": "\\",
              "shift": "|"
            }
          ],
          [
            {
              "default": "Caps"
            },
            {
              "default": "a",
              "shift": "A",
              "caps": "A"
            },
            {
              "default": "s",
              "shift": "S",
              "caps": "S"
            },
            {
              "default": "d",
              "shift": "D",
              "caps": "D"
            },
            {
              "default": "f",
              "shift": "F",
              "caps": "F"
            },
            {
              "default": "g",
              "shift": "G",
              "caps": "G"
            },
            {
              "default": "h",
              "shift": "H",
              "caps": "H"
            },
            {
              "default": "j",
              "shift": "J",
              "caps": "J"
            },
            {
              "default": "k",
              "shift": "K",
              "caps": "K"
            },
            {
              "default": "l",
              "shift": "L",
              "caps": "L"
            },
            {
              "default": ";",
              "shift": ":"
            },
            {
              "default": "'",
              "shift": "\""
            },
            {
              "default": "Enter"
            }
          ],
          [
            {
              "default": "Shift"
            },
            {
              "default": "z",
              "shift": "Z",
              "caps": "Z"
            },
            {
              "default": "x",
              "shift": "X",
              "caps": "X"
            },
            {
              "default": "c",
              "shift": "C",
              "caps": "C"
            },
            {
              "default": "v",
              "shift": "V",
              "caps": "V"
            },
            {
              "default": "b",
              "shift": "B",
              "caps": "B"
            },
            {
              "default": "n",
              "shift": "N",
              "caps": "N"
            },
            {
              "default": "m",
              "shift": "M",
              "caps": "M"
            },
            {
              "default": ",",
              "shift": "<"
            },
            {
              "default": ".",
              "shift": ">"
            },
            {
              "default": "/",
              "shift": "?"
            },
            {
              "default": "Shift"
            }
          ],
          [
            {
              "default": "Ctrl"
            },
            {
              "default": "Win"
            },
            {
              "default": "Alt"
            },
            {
              "default": "Space"
            },
            {
              "default": "Alt"
            },
            {
              "default": "Men"
            },
            {
              "default": "Ctrl"
            },
            {
              "default": "&larr;"
            },
            {
              "default": "&uarr;"
            },
            {
              "default": "&darr;"
            },
            {
              "default": "&rarr;"
            }
          ]
      ];

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
