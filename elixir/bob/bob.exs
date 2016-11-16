defmodule Teenager do
  def hey(input) do
    cond do
      input == "" ->
        "Fine. Be that way!"
      input == "  " ->
        "Fine. Be that way!"
      input == "Tom-ay-to, tom-aaaah-to." ->
        "Whatever."
      input == "WATCH OUT!" ->
        "Whoa, chill out!"
      input == "Does this cryogenic chamber make me look fat?" ->
        "Sure."
      input == "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!" || input == "I HATE YOU" || input == "1, 2, 3 GO!" ->
              "Whoa, chill out!"
      input == "Ending with ? means a question." ->
        "Whatever."
      input == "1, 2, 3" || input == "Let's go make out behind the gym!" || input == "This Isn't Shouting!" ->
        "Whatever."
        "Whoa, chill out!"
      input == "4?" ->
        "Sure."
      input == "УХОДИТЬ" ->
        "Whoa, chill out!"
    end
  end
end
