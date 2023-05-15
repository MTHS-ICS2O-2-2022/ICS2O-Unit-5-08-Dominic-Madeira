// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"


/**
 * This function uses a selection component from https://github.com/CreativeIT/getmdl-select
 */

function myButtonClicked() {

  // input
  const numberA = parseInt(document.getElementById('number-a').value)
  const numberB = parseInt(document.getElementById('number-b').value)
  let remainder = numberA
  let output = 0

  // process

  while (remainder >= numberB) {
    remainder -= numberB
    output++
  }
  if (remainder > 0) {
    document.getElementById('answer').innerHTML = numberA + " ÷ " + numberB + " = " + output + " with a remainder of " + remainder + "."
  } else {
    document.getElementById('answer').innerHTML = numberA + " ÷ " + numberB + " = " + output + "."
  }
}