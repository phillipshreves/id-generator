# Human-readable ID Generator

To generate the IDs, it takes input in the form of JSON arrays containing english words and outputs them to a text list. During this process, it appends characters to each word to reach a specified total length.


## Why?

The intention is to have IDs that can easily be read, shared, and input easily into mobile devices as a way to create shared sessions via unique IDs.


## How?

For this specific use case, we include 3 lists, each containing different lengths of words(4, 5, and 6):

For the 6 character words, it simply appends them to the list. 

For the 5 character words, it intakes each word and iterates over it, apending the characters A-Z to the end of it, so that one ID becomes 26 IDs(TESTSA, TESTSB, ...TESTSY, TESTSZ). 

Fo the 4 character words, it intakes each word and iterates over it similar to the 5 character words, however, now that we are adding two letters, we now transform each ID into over 600 unique IDs(TESTAA, TESTAB, ...TESTZY, TESTZZ).

From approximately 10,000 unique IDs, we've created over 1.5 million IDs.

We then randomize the list of IDs so that they are not sequential and cannot be guessed.





<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><a property="dct:title" rel="cc:attributionURL" href="https://github.com/phillipshreves/id-generator">id-generator</a> by <a rel="cc:attributionURL dct:creator" property="cc:attributionName" href="https://github.com/phillipshreves">Phillip Shreves</a> is licensed under <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">CC BY-NC-SA 4.0<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/nc.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/sa.svg?ref=chooser-v1"></a></p>
