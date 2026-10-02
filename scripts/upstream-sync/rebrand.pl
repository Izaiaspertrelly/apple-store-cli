#!/usr/bin/perl -pi
# Rebrand upstream (rorkai/rudrankriyam App-Store-Connect-CLI) -> Izaiaspertrelly/apple-store-cli.
# Mirrors the mapping used in the original 5.1.0 fork. Idempotent.
s#gitlab\.com/rudrankriyam/asc-ci-components/(\w+)\@main#github.com/Izaiaspertrelly/apple-store-cli/$1\@main#g;
s#https://gitlab\.com/rudrankriyam/asc-ci-components#https://github.com/Izaiaspertrelly/apple-store-cli#g;
s#(https://)?github\.com/rudrankriyam/steps-setup-asc#${1}github.com/Izaiaspertrelly/setup-asc#g;
s#rudrankriyam/setup-asc#Izaiaspertrelly/setup-asc#g;
s#https://github\.com/rudrankriyam/(asc-orb|asc-ci-components)#https://github.com/Izaiaspertrelly/apple-store-cli#g;
s#rudrankriyam/asc\@#pertrelly/asc\@#g;
s#rudrankriyam/koubou#bitomule/koubou#g;
s#com\.rudrankriyam\.#com.example.#g;
s#(?:rorkai|rudrankriyam)/App-Store-Connect-CLI(?!-skills)#Izaiaspertrelly/apple-store-cli#gi;
s#rorkai/user-workflows#Izaiaspertrelly/user-workflows#g;
s#Rorkai\.ASC#Pertrelly.ASC#g;
s#manifests/r/Rorkai/ASC#manifests/p/Pertrelly/ASC#g;
s#rorkai-asc-#pertrelly-asc-#g;
s#App-Store-Connect-CLI(?!-skills)#apple-store-cli#g;
s#https://x\.com/rudrankriyam#https://github.com/Izaiaspertrelly#g;
s#https://github\.com/rudrankriyam\b#https://github.com/Izaiaspertrelly#g;
s#\bApp Store Connect CLI\b#Apple Store CLI#g;
s#Connect from Rork\.#Connect.#g;
s#com\.rorkai\.asc#com.pertrelly.asc#g;
s#com\.rudrank\.#com.example.#g;
s#rudrankriyam\@gmail\.com#pertrelly\@echohub.ai#g;
s#release-bot\@rork\.com#release-bot\@users.noreply.github.com#g;
