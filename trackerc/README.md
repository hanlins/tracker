INTRODUCTION
============

This test is relative 'fragile', and follow the steps listed, you can replay
the tests I did.

TESTING ENVIRONMENT
===================

* Browser : Google Chrome Canary
* Setting : Third-party cookie blocking `enabled`,
            Popup window blocking `disabled`

TESTING STEPS
=============

* Clear browser history (cookie for domain `trackerc.com`)
* Go to `mylab.com` -> `Tracker type C`, ==> create a popup and set cookie,
                                             close the popup or leave it.
* Refresh the page `Tracker type C` ( should be redirected to
  `http://trackerc.com/c/base.html` )

Now on the tracker's monitor, you can see that the visiting has been recorded.
All following visits that embeds this tracker will be recorded.
