import {
  LitElement,
  html,
} from "https://cdn.jsdelivr.net/gh/lit/dist@2/all/lit-all.min.js";
import { create, cssomSheet } from "https://cdn.skypack.dev/twind";
import config from "/lib/twind.config.mjs";

// 1. Create separate CSSStyleSheet
const sheet = cssomSheet({ target: new CSSStyleSheet() });

// 2. Use that to create an own twind instance
const { tw } = create({ ...config, sheet });

export class SimpleGreeting extends LitElement {
  static properties = {
    count: 0,
  };

  static styles = [sheet.target];

  constructor() {
    super();
    // Declare reactive properties
    this.count = 0;
    setInterval(() => {
      this.count++;
    }, 500);
  }

  // Render the UI as a function of component state
  render() {
    return html`
      <p>Current count: ${this.count}</p>
      <div>haha!</div>
      <button
        @click=${() => {
          this.count = 0;
        }}
        class="${tw`bg-red-500`}"
      >
        reset
      </button>
    `;
  }
}
customElements.define("simple-greeting", SimpleGreeting);
