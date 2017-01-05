package kendo

type AutoComplete struct {
  // Configures the opening and closing animations of the suggestion popup.
  Animation     Animation
  // If set to true, the widget automatically adjusts the width of the popup element and does not wrap up the item label.
  // Note: Virtualized list doesn't support the auto-width functionality.
  AutoWidth     bool
}