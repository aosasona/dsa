package algorithms

import "testing"

const haystack = [9]string{"Harley", "Bruce", "Clark", "Diana", "Luke", "Lois", "Kara", "Tony", "Pietro"}

func Test_LinearSearchPasses(t *testing.Testing) {
  var (
    needle = "Lois" 
  )

  if found := LinearSearch(needle, haystack); !found {
    t.Errorf("Failed to find %s, expected `true`, got `%s`", needle, found)
  }
}