import * as fs from 'fs';
import * as os from 'os';
import * as path from 'path';

import * as context from '../src/context';

const pgp = `-----BEGIN PGP PRIVATE KEY BLOCK-----

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
fLEIDM4LAqCGRN2Mx5OLHuM0ZkZnF14=
=WbwB
-----END PGP PRIVATE KEY BLOCK-----`;

jest.spyOn(context, 'defaultContext').mockImplementation((): string => {
  return 'https://github.com/docker/build-push-action.git#refs/heads/test-jest';
});

jest.spyOn(context, 'tmpDir').mockImplementation((): string => {
  const tmpDir = path.join('/tmp/.docker-build-push-jest').split(path.sep).join(path.posix.sep);
  if (!fs.existsSync(tmpDir)) {
    fs.mkdirSync(tmpDir, {recursive: true});
  }
  return tmpDir;
});

jest.spyOn(context, 'tmpNameSync').mockImplementation((): string => {
  return path.join('/tmp/.docker-build-push-jest', '.tmpname-jest').split(path.sep).join(path.posix.sep);
});

describe('getArgs', () => {
  beforeEach(() => {
    process.env = Object.keys(process.env).reduce((object, key) => {
      if (!key.startsWith('INPUT_')) {
        object[key] = process.env[key];
      }
      return object;
    }, {});
  });

  // prettier-ignore
  test.each([
    [
      '0.4.1',
      new Map<string, string>([
        ['context', '.'],
      ]),
      [
        'buildx',
        'build',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        '.'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['build-args', 'MY_ARG=val1,val2,val3\nARG=val'],
      ]),
      [
        'buildx',
        'build',
        '--build-arg', 'MY_ARG=val1,val2,val3',
        '--build-arg', 'ARG=val',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        'https://github.com/docker/build-push-action.git#refs/heads/test-jest'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['tags', 'name/app:7.4, name/app:latest'],
      ]),
      [
        'buildx',
        'build',
        '--tag', 'name/app:7.4',
        '--tag', 'name/app:latest',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        'https://github.com/docker/build-push-action.git#refs/heads/test-jest'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['context', '.'],
        ['labels', 'org.opencontainers.image.title=buildkit\norg.opencontainers.image.description=concurrent, cache-efficient, and Dockerfile-agnostic builder toolkit'],
        ['outputs', 'type=local,dest=./release-out']
      ]),
      [
        'buildx',
        'build',
        '--label', 'org.opencontainers.image.title=buildkit',
        '--label', 'org.opencontainers.image.description=concurrent, cache-efficient, and Dockerfile-agnostic builder toolkit',
        '--output', 'type=local,dest=./release-out',
        '.'
      ]
    ],
    [
      '0.4.1',
      new Map<string, string>([
        ['context', '.'],
        ['platforms', 'linux/amd64,linux/arm64']
      ]),
      [
        'buildx',
        'build',
        '--platform', 'linux/amd64,linux/arm64',
        '.'
      ]
    ],
    [
      '0.4.1',
      new Map<string, string>([
        ['context', '.']
      ]),
      [
        'buildx',
        'build',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        '.'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['context', '.'],
        ['secrets', 'GIT_AUTH_TOKEN=abcdefghijklmno=0123456789'],
      ]),
      [
        'buildx',
        'build',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        '--secret', 'id=GIT_AUTH_TOKEN,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '.'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['github-token', 'abcdefghijklmno0123456789'],
        ['outputs', '.']
      ]),
      [
        'buildx',
        'build',
        '--output', '.',
        '--secret', 'id=GIT_AUTH_TOKEN,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        'https://github.com/docker/build-push-action.git#refs/heads/test-jest'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['context', 'https://github.com/docker/build-push-action.git#refs/heads/master'],
        ['tag', 'localhost:5000/name/app:latest'],
        ['platforms', 'linux/amd64,linux/arm64'],
        ['secrets', 'GIT_AUTH_TOKEN=abcdefghijklmno=0123456789'],
        ['file', './test/Dockerfile'],
        ['builder', 'builder-git-context-2'],
        ['push', 'true']
      ]),
      [
        'buildx',
        'build',
        '--platform', 'linux/amd64,linux/arm64',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        '--secret', 'id=GIT_AUTH_TOKEN,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--file', './test/Dockerfile',
        '--builder', 'builder-git-context-2',
        '--push',
        'https://github.com/docker/build-push-action.git#refs/heads/master'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['context', 'https://github.com/docker/build-push-action.git#refs/heads/master'],
        ['tag', 'localhost:5000/name/app:latest'],
        ['platforms', 'linux/amd64,linux/arm64'],
        ['secrets', `GIT_AUTH_TOKEN=abcdefghi,jklmno=0123456789
"MYSECRET=aaaaaaaa
bbbbbbb
ccccccccc"
FOO=bar
"EMPTYLINE=aaaa

bbbb
ccc"`],
        ['file', './test/Dockerfile'],
        ['builder', 'builder-git-context-2'],
        ['push', 'true']
      ]),
      [
        'buildx',
        'build',
        '--platform', 'linux/amd64,linux/arm64',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        '--secret', 'id=GIT_AUTH_TOKEN,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--secret', 'id=MYSECRET,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--secret', 'id=FOO,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--secret', 'id=EMPTYLINE,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--file', './test/Dockerfile',
        '--builder', 'builder-git-context-2',
        '--push',
        'https://github.com/docker/build-push-action.git#refs/heads/master'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['context', 'https://github.com/docker/build-push-action.git#refs/heads/master'],
        ['tag', 'localhost:5000/name/app:latest'],
        ['platforms', 'linux/amd64,linux/arm64'],
        ['secrets', `GIT_AUTH_TOKEN=abcdefghi,jklmno=0123456789
MYSECRET=aaaaaaaa
bbbbbbb
ccccccccc
FOO=bar
EMPTYLINE=aaaa

bbbb
ccc`],
        ['file', './test/Dockerfile'],
        ['builder', 'builder-git-context-2'],
        ['push', 'true']
      ]),
      [
        'buildx',
        'build',
        '--platform', 'linux/amd64,linux/arm64',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        '--secret', 'id=GIT_AUTH_TOKEN,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--secret', 'id=MYSECRET,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--secret', 'id=FOO,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--secret', 'id=EMPTYLINE,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--file', './test/Dockerfile',
        '--builder', 'builder-git-context-2',
        '--push',
        'https://github.com/docker/build-push-action.git#refs/heads/master'
      ]
    ],
    [
      '0.5.1',
      new Map<string, string>([
        ['context', 'https://github.com/docker/build-push-action.git#refs/heads/master'],
        ['tag', 'localhost:5000/name/app:latest'],
        ['secret-files', `MY_SECRET=${path.join(__dirname, 'fixtures', 'secret.txt').split(path.sep).join(path.posix.sep)}`],
        ['file', './test/Dockerfile'],
        ['builder', 'builder-git-context-2'],
        ['network', 'host'],
        ['push', 'true']
      ]),
      [
        'buildx',
        'build',
        '--iidfile', '/tmp/.docker-build-push-jest/iidfile',
        '--secret', 'id=MY_SECRET,src=/tmp/.docker-build-push-jest/.tmpname-jest',
        '--file', './test/Dockerfile',
        '--builder', 'builder-git-context-2',
        '--network', 'host',
        '--push',
        'https://github.com/docker/build-push-action.git#refs/heads/master'
      ]
    ],
    [
      '0.4.2',
      new Map<string, string>([
        ['context', '.'],
        ['labels', 'org.opencontainers.image.title=filter_results_top_n\norg.opencontainers.image.description=Reference implementation of operation "filter results (top-n)"'],
        ['outputs', 'type=local,dest=./release-out']
      ]),
      [
        'buildx',
        'build',
        '--label', 'org.opencontainers.image.title=filter_results_top_n',
        '--label', 'org.opencontainers.image.description=Reference implementation of operation "filter results (top-n)"',
        '--output', 'type=local,dest=./release-out',
        '.'
      ]
    ]
  ])(
    'given %p with %p as inputs, returns %p',
    async (buildxVersion: string, inputs: Map<string, any>, expected: Array<string>) => {
      await inputs.forEach((value: string, name: string) => {
        setInput(name, value);
      });
      const defContext = context.defaultContext();
      const inp = await context.getInputs(defContext);
      console.log(inp);
      const res = await context.getArgs(inp, defContext, buildxVersion);
      console.log(res);
      expect(res).toEqual(expected);
    }
  );
});

describe('getInputList', () => {
  it('single line correctly', async () => {
    await setInput('foo', 'bar');
    const res = await context.getInputList('foo');
    console.log(res);
    expect(res).toEqual(['bar']);
  });

  it('multiline correctly', async () => {
    setInput('foo', 'bar\nbaz');
    const res = await context.getInputList('foo');
    console.log(res);
    expect(res).toEqual(['bar', 'baz']);
  });

  it('empty lines correctly', async () => {
    setInput('foo', 'bar\n\nbaz');
    const res = await context.getInputList('foo');
    console.log(res);
    expect(res).toEqual(['bar', 'baz']);
  });

  it('comma correctly', async () => {
    setInput('foo', 'bar,baz');
    const res = await context.getInputList('foo');
    console.log(res);
    expect(res).toEqual(['bar', 'baz']);
  });

  it('empty result correctly', async () => {
    setInput('foo', 'bar,baz,');
    const res = await context.getInputList('foo');
    console.log(res);
    expect(res).toEqual(['bar', 'baz']);
  });

  it('different new lines correctly', async () => {
    setInput('foo', 'bar\r\nbaz');
    const res = await context.getInputList('foo');
    console.log(res);
    expect(res).toEqual(['bar', 'baz']);
  });

  it('different new lines and comma correctly', async () => {
    setInput('foo', 'bar\r\nbaz,bat');
    const res = await context.getInputList('foo');
    console.log(res);
    expect(res).toEqual(['bar', 'baz', 'bat']);
  });

  it('multiline and ignoring comma correctly', async () => {
    setInput('cache-from', 'user/app:cache\ntype=local,src=path/to/dir');
    const res = await context.getInputList('cache-from', true);
    console.log(res);
    expect(res).toEqual(['user/app:cache', 'type=local,src=path/to/dir']);
  });

  it('different new lines and ignoring comma correctly', async () => {
    setInput('cache-from', 'user/app:cache\r\ntype=local,src=path/to/dir');
    const res = await context.getInputList('cache-from', true);
    console.log(res);
    expect(res).toEqual(['user/app:cache', 'type=local,src=path/to/dir']);
  });

  it('multiline values', async () => {
    setInput(
      'secrets',
      `GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789
"MYSECRET=aaaaaaaa
bbbbbbb
ccccccccc"
FOO=bar`
    );
    const res = await context.getInputList('secrets', true);
    console.log(res);
    expect(res).toEqual([
      'GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789',
      `MYSECRET=aaaaaaaa
bbbbbbb
ccccccccc`,
      'FOO=bar'
    ]);
  });

  it('multiline values with empty lines', async () => {
    setInput(
      'secrets',
      `GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789
"MYSECRET=aaaaaaaa
bbbbbbb
ccccccccc"
FOO=bar
"EMPTYLINE=aaaa

bbbb
ccc"`
    );
    const res = await context.getInputList('secrets', true);
    console.log(res);
    expect(res).toEqual([
      'GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789',
      `MYSECRET=aaaaaaaa
bbbbbbb
ccccccccc`,
      'FOO=bar',
      `EMPTYLINE=aaaa

bbbb
ccc`
    ]);
  });

  it('multiline values without quotes', async () => {
    setInput(
      'secrets',
      `GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789
MYSECRET=aaaaaaaa
bbbbbbb
ccccccccc
FOO=bar`
    );
    const res = await context.getInputList('secrets', true);
    console.log(res);
    expect(res).toEqual([
      'GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789',
      'MYSECRET=aaaaaaaa',
      'bbbbbbb',
      'ccccccccc',
      'FOO=bar'
    ]);
  });

  it('large multiline values', async () => {
    setInput(
      'secrets',
      `"GPG_KEY=${pgp}"
FOO=bar`
    );
    const res = await context.getInputList('secrets', true);
    console.log(res);
    expect(res).toEqual([`GPG_KEY=${pgp}`, 'FOO=bar']);
  });

  it('multiline values escape quotes', async () => {
    setInput(
      'secrets',
      `GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789
"MYSECRET=aaaaaaaa
bbbb""bbb
ccccccccc"
FOO=bar`
    );
    const res = await context.getInputList('secrets', true);
    console.log(res);
    expect(res).toEqual([
      'GIT_AUTH_TOKEN=abcdefgh,ijklmno=0123456789',
      `MYSECRET=aaaaaaaa
bbbb\"bbb
ccccccccc`,
      'FOO=bar'
    ]);
  });
});

describe('asyncForEach', () => {
  it('executes async tasks sequentially', async () => {
    const testValues = [1, 2, 3, 4, 5];
    const results: number[] = [];

    await context.asyncForEach(testValues, async value => {
      results.push(value);
    });

    expect(results).toEqual(testValues);
  });
});

describe('setOutput', () => {
  beforeEach(() => {
    process.stdout.write = jest.fn();
  });

  it('setOutput produces the correct command', () => {
    context.setOutput('some output', 'some value');
    assertWriteCalls([`::set-output name=some output::some value${os.EOL}`]);
  });

  it('setOutput handles bools', () => {
    context.setOutput('some output', false);
    assertWriteCalls([`::set-output name=some output::false${os.EOL}`]);
  });

  it('setOutput handles numbers', () => {
    context.setOutput('some output', 1.01);
    assertWriteCalls([`::set-output name=some output::1.01${os.EOL}`]);
  });
});

// See: https://github.com/actions/toolkit/blob/master/packages/core/src/core.ts#L67
function getInputName(name: string): string {
  return `INPUT_${name.replace(/ /g, '_').toUpperCase()}`;
}

function setInput(name: string, value: string): void {
  process.env[getInputName(name)] = value;
}

// Assert that process.stdout.write calls called only with the given arguments.
function assertWriteCalls(calls: string[]): void {
  expect(process.stdout.write).toHaveBeenCalledTimes(calls.length);
  for (let i = 0; i < calls.length; i++) {
    expect(process.stdout.write).toHaveBeenNthCalledWith(i + 1, calls[i]);
  }
}
