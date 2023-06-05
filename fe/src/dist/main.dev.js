"use strict";

function _typeof(obj) { if (typeof Symbol === "function" && typeof Symbol.iterator === "symbol") { _typeof = function _typeof(obj) { return typeof obj; }; } else { _typeof = function _typeof(obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; }; } return _typeof(obj); }

var _vue = require("vue");

var _App = _interopRequireDefault(require("./App.vue"));

var _router = _interopRequireDefault(require("./router"));

require("vuetify/styles");

var _vuetify = require("vuetify");

var components = _interopRequireWildcard(require("vuetify/components"));

var directives = _interopRequireWildcard(require("vuetify/directives"));

var _eventbus = _interopRequireDefault(require("./eventbus"));

var _vueSimpleWebsocket = _interopRequireDefault(require("vue-simple-websocket"));

function _getRequireWildcardCache() { if (typeof WeakMap !== "function") return null; var cache = new WeakMap(); _getRequireWildcardCache = function _getRequireWildcardCache() { return cache; }; return cache; }

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } if (obj === null || _typeof(obj) !== "object" && typeof obj !== "function") { return { "default": obj }; } var cache = _getRequireWildcardCache(); if (cache && cache.has(obj)) { return cache.get(obj); } var newObj = {}; var hasPropertyDescriptor = Object.defineProperty && Object.getOwnPropertyDescriptor; for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) { var desc = hasPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : null; if (desc && (desc.get || desc.set)) { Object.defineProperty(newObj, key, desc); } else { newObj[key] = obj[key]; } } } newObj["default"] = obj; if (cache) { cache.set(obj, newObj); } return newObj; }

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

// Vuetify
var vuetify = (0, _vuetify.createVuetify)({
  components: components,
  directives: directives
});
var app = (0, _vue.createApp)(_App["default"]);
app.use(_vueSimpleWebsocket["default"], "ws://localhost:3000/ws", {
  reconnectEnabled: true,
  reconnectInterval: 1000
});
app.use(_router["default"]);
app.use(vuetify);
app.mount('#app');