# delete oldest tweet
Delete's the oldest tweet for the authenticated Twitter user based on the tweet displayed at https://discover.twitter.com/first-tweet

It then loops to delete the next oldest tweet.

In effect, this tries to **delete all tweets** starting with the oldest.

How this differs from other "delete all tweets" things is that it scrapes the above page to find the tweet id of the oldest tweet and then deletes that. This allows it to truly delete all tweets, rather than just the tweets that are returned by searches (which are limited to a recent window, and so a lot of tweets will never be deleted).

It may be that Twitter stop returning the last tweet, there's not a lot I can do about that.

You need to be registered as a Twitter developer, and to create an app and then auth your user against the app: https://apps.twitter.com

This is MIT licence, no warranty or support implied or offered, use at own risk, and it will delete your tweets and this is not reversible.

Alternatives for more recent tweets and in bulk: http://www.tweetdelete.net/
