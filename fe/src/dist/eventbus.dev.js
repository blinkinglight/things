"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _instance = _interopRequireDefault(require("tiny-emitter/instance"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

// eventBus.js
var _default = {
  $on: function $on() {
    return _instance["default"].on.apply(_instance["default"], arguments);
  },
  $once: function $once() {
    return _instance["default"].once.apply(_instance["default"], arguments);
  },
  $off: function $off() {
    return _instance["default"].off.apply(_instance["default"], arguments);
  },
  $emit: function $emit() {
    return _instance["default"].emit.apply(_instance["default"], arguments);
  }
};
exports["default"] = _default;