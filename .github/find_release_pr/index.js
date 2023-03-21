// This script searches the current repository for a release PR, extracts the
// version and branch from it, and sets them as outputs.
const core = require('@actions/core');
const exec = require('@actions/exec');

const regex = new RegExp('^Release v((0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?)$', 'gm')

async function main() {
    try {
        await mainE();
    } catch (error) {
        core.setFailed(error.message);
    }
}

async function mainE() {
    const number = await findReleasePRNumber();
    const pr = await viewPR(number);
    core.info('Found release PR: ' + pr.title);

    const version = extractVersionFromTitle(pr.title);
    core.info('Extracted version: ' + version);
    const branch = extractBranchFromHeadRefName(pr.headRefName);
    core.info('Extracted branch: ' + branch);

    core.setOutput('version', version);
    core.setOutput('branch', branch);
}

async function findReleasePRNumber() {
    const { code, output, error } = await gh(['pr', 'list', '--repo', 'giantswarm/helm-values-gen', '--state', 'open', '--json', 'title,number'])
    if (code != 0) {
        throw new Error(error);
    }
    const prs = JSON.parse(output.join(''));
    const releasePRs = prs.filter(pr => pr.title.match(regex));

    if (releasePRs.length !== 1) {
        throw new Error('Unexpected number of release PRs found: ' + releasePRs.length);
    }

    const releasePR = releasePRs[0];
    const number = releasePR.number;

    return number;
}

async function gh(args) {
    core.info('Running: gh ' + args.join(' '));

    const output = [];
    const error = [];

    const options = { silent: true }
    options.listeners = {
        stdout: (data) => {
            output.push(data.toString());
            core.debug(data.toString());
        },
        stderr: (data) => {
            error.push(data.toString());
            core.debug(data.toString());
        }
    };
    code = await exec.exec('bash', ['-c', 'gh ' + args.join(' ')], options);
    return { code: code, output, error };
}

function extractVersionFromTitle(title) {
    const match = regex.exec(title);
    const version = match[1];
    return version;
}

function extractBranchFromHeadRefName(headRefName) {
    const branch = headRefName.replace('refs/heads/', '');
    return branch;
}

async function viewPR(number) {
    const { code, output, error } = await gh(['pr', 'view', number, '--json', 'title,headRefName'])
    if (code != 0) {
        throw new Error(error);
    }
    const pr = JSON.parse(output.join(''));
    return pr;
}

main();
