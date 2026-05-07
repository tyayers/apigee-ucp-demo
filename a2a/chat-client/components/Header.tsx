/*
 * Copyright 2026 UCP Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

interface HeaderProps {
  logoUrl: string;
  title: string;
  mode: "freemium" | "pro";
  onModeChange: (mode: "freemium" | "pro") => void;
}

function Header({ logoUrl, title, mode, onModeChange }: HeaderProps) {
  return (
    <header className="bg-white shadow-sm p-4 border-b border-gray-200 flex-shrink-0 flex justify-between items-center">
      <div className="flex items-center">
        <img src={logoUrl} alt={title} className="h-8 mr-3" />
        <span className="text-xl font-bold text-gray-800">{title}</span>
      </div>
      <div className="flex items-center space-x-2">
        <span
          className={`text-sm font-medium ${mode === "freemium" ? "text-blue-600" : "text-gray-400"}`}
        >
          Freemium
        </span>
        <button
          onClick={() => onModeChange(mode === "freemium" ? "pro" : "freemium")}
          className="relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none bg-gray-200"
          role="switch"
          aria-checked={mode === "pro"}
        >
          <span
            aria-hidden="true"
            className={`pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out ${mode === "pro" ? "translate-x-5" : "translate-x-0"}`}
          />
        </button>
        <span
          className={`text-sm font-medium ${mode === "pro" ? "text-blue-600" : "text-gray-400"}`}
        >
          Pro
        </span>
      </div>
    </header>
  );
}

export default Header;
